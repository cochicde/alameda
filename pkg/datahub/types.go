package datahub

import (
	Entities "github.com/containers-ai/alameda/datahub/pkg/entities"
	"github.com/containers-ai/alameda/datahub/pkg/formatconversion/responses/enumconv"
	DBCommon "github.com/containers-ai/alameda/pkg/database/common"
	"github.com/containers-ai/alameda/pkg/utils"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/data"
	"github.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas"
	"github.com/golang/protobuf/ptypes/duration"
	"reflect"
	"strconv"
	"time"
)

type Option struct {
	Entity interface{}
	Fields []string
}

type TimeRange struct {
	StartTime *time.Time
	EndTime   *time.Time
	Order     DBCommon.Order
	Limit     uint64
	Step      int
}

type Function struct {
	Type              DBCommon.FunctionType
	Target            string
	Unit              string
	Number            int64
}

func NewWriteData(entities interface{}, fields []string) *data.WriteData {
	writeData := data.WriteData{}
	measurement := ExtractField(entities, "Measurement")
	writeData.Measurement = measurement.Tag.Get("name")
	writeData.MetricType = MetricTypeValue[measurement.Tag.Get("metric")]
	writeData.ResourceBoundary = ResourceBoundaryValue[measurement.Tag.Get("boundary")]
	writeData.ResourceQuota = ResourceQuotaValue[measurement.Tag.Get("quota")]
	writeData.Columns = NewColumns(entities, fields)
	writeData.Rows = NewRows(entities, writeData.Columns)
	return &writeData
}

func NewReadData(entities interface{}, fields []string, timeRange *TimeRange, function *Function, opts ...Option) *data.ReadData {
	readData := data.ReadData{}
	selects := make([]string, 0)
	measurement := ExtractField(entities, "Measurement")
	readData.Measurement = measurement.Tag.Get("name")
	readData.MetricType = MetricTypeValue[measurement.Tag.Get("metric")]
	readData.ResourceBoundary = ResourceBoundaryValue[measurement.Tag.Get("boundary")]
	readData.ResourceQuota = ResourceQuotaValue[measurement.Tag.Get("quota")]
	entityType := reflect.TypeOf(entities).Elem().Elem()
	for _, field := range fields {
		f, _ := entityType.FieldByName(field)
		selects = append(selects, f.Tag.Get("json"))
	}
	readData.QueryCondition = NewQueryCondition(selects, timeRange, function, opts...)
	return &readData
}

func NewDeleteData(entities interface{}, opts ...Option) *data.DeleteData {
	deleteData := data.DeleteData{}
	measurement := ExtractField(entities, "Measurement")
	deleteData.Measurement = measurement.Tag.Get("name")
	deleteData.MetricType = MetricTypeValue[measurement.Tag.Get("metric")]
	deleteData.ResourceBoundary = ResourceBoundaryValue[measurement.Tag.Get("boundary")]
	deleteData.ResourceQuota = ResourceQuotaValue[measurement.Tag.Get("quota")]
	deleteData.QueryCondition = NewQueryCondition(nil, nil, nil, opts...)
	return &deleteData
}

func NewSchemaMeta(entities interface{}) *schemas.SchemaMeta {
	schemaMeta := schemas.SchemaMeta{}
	datahubEntity := ExtractField(entities, "DatahubEntity")
	schemaMeta.Scope = ScopeValue[datahubEntity.Tag.Get("scope")]
	schemaMeta.Category = datahubEntity.Tag.Get("category")
	schemaMeta.Type = datahubEntity.Tag.Get("type")
	return &schemaMeta
}

func NewColumns(entities interface{}, fields []string) []string {
	columns := make([]string, 0)

	entityType := reflect.TypeOf(entities).Elem().Elem()
	fieldNames := fields

	// If fields is empty which means to iterate through all the fields of the entity
	if len(fieldNames) == 0 {
		// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
		for i := Entities.ColumnStartIndex + 1; i < entityType.NumField(); i++ {
			fieldNames = append(fieldNames, entityType.Field(i).Name)
		}
	} else {
		// Add influx tags if it was not added
		// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
		for i := Entities.ColumnStartIndex + 1; i < entityType.NumField(); i++ {
			fieldType := entityType.Field(i)
			if fieldType.Tag.Get("column") == "tag" {
				if !utils.SliceContains(fieldNames, fieldType.Name) {
					fieldNames = append(fieldNames, fieldType.Name)
				}
			}
		}
	}

	// Read the tag of field to generate the column list
	for _, field := range fieldNames {
		// Index is started at 3 to skip DatahubEntity, Measurement and Time fields
		for i := Entities.ColumnStartIndex + 1; i < entityType.NumField(); i++ {
			fieldType := entityType.Field(i)
			if fieldType.Name == field {
				if !utils.SliceContains(columns, fieldType.Tag.Get("json")) {
					columns = append(columns, fieldType.Tag.Get("json"))
				}
				break
			}
		}
	}

	return columns
}

func NewRows(entities interface{}, columns []string) []*common.Row {
	rows := make([]*common.Row, 0)

	values := reflect.ValueOf(entities).Elem()
	for i := 0; i < values.Len(); i++ {
		rows = append(rows, NewRow(values.Index(i), columns))
	}

	return rows
}

func NewRow(value reflect.Value, columns []string) *common.Row {
	row := common.Row{}
	row.Time = utils.TimestampProto(value.FieldByName("Time").Interface().(*time.Time))

	for _, column := range columns {
		// Index is started at 2 to skip DatahubEntity and Measurement fields
		for i := Entities.ColumnStartIndex; i < value.NumField(); i++ {
			fieldValue := value.Field(i)
			fieldType := value.Type().Field(i)
			value := ""
			if column == fieldType.Tag.Get("json") {
				switch fieldValue.Kind() {
				case reflect.Bool:
					value = strconv.FormatBool(reflect.ValueOf(fieldValue.Interface()).Bool())
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					value = strconv.FormatInt(reflect.ValueOf(fieldValue.Interface()).Int(), 10)
				case reflect.Float32, reflect.Float64:
					value = strconv.FormatFloat(reflect.ValueOf(fieldValue.Interface()).Float(), 'f', -1, 64)
				case reflect.String:
					value = reflect.ValueOf(fieldValue.Interface()).String()
				case reflect.Ptr:
					// Since only Time field is pointer type, skip this type currently
					continue
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}

				row.Values = append(row.Values, value)
				break
			}
		}
	}

	return &row
}

func NewQueryCondition(selects []string, timeRange *TimeRange, function *Function, opts ...Option) *common.QueryCondition {
	if len(selects) == 0 && timeRange == nil && function == nil && len(opts) == 0 {
		return nil
	}

	queryCondition := common.QueryCondition{}
	queryCondition.TimeRange = NewTimeRange(timeRange)
	queryCondition.Function = NewFunction(function)
	queryCondition.Selects = selects
	if timeRange != nil {
		if timeRange.Order != 0 {
			queryCondition.Order = enumconv.QueryConditionOrderNameMap[timeRange.Order]
		}
		if timeRange.Limit != 0 {
			queryCondition.Limit = timeRange.Limit
		}
	}
	for _, opt := range opts {
		queryCondition.WhereCondition = append(queryCondition.WhereCondition, NewCondition(opt))
	}

	return &queryCondition
}

func NewCondition(opt Option) *common.Condition {
	condition := common.Condition{}

	values := reflect.ValueOf(&opt.Entity).Elem().Elem()
	for _, field := range opt.Fields {
		for i := 0; i < values.NumField(); i++ {
			fieldType := values.Type().Field(i)
			if field == fieldType.Name {
				fieldValue := values.FieldByName(field)
				value := ""

				switch fieldValue.Interface().(type) {
				case bool:
					value = strconv.FormatBool(reflect.ValueOf(fieldValue.Interface()).Bool())
				case int, int8, int16, int32, int64:
					value = strconv.FormatInt(reflect.ValueOf(fieldValue.Interface()).Int(), 10)
				case float32, float64:
					value = strconv.FormatFloat(reflect.ValueOf(fieldValue.Interface()).Float(), 'f', -1, 64)
				case string:
					value = reflect.ValueOf(fieldValue.Interface()).String()
				case *time.Time:
					value = fieldValue.Interface().(*time.Time).Format(time.RFC3339)
				default:
					scope.Errorf("field type(%s) not supported", fieldValue.Kind().String())
				}

				condition.Keys = append(condition.Keys, fieldType.Tag.Get("json"))
				condition.Values = append(condition.Values, value)
				condition.Operators = append(condition.Operators, "=")
				break
			}
		}
	}

	return &condition
}

func NewTimeRange(tr *TimeRange) *common.TimeRange {
	if tr != nil {
		timeRange := common.TimeRange{}
		if tr.StartTime != nil {
			timeRange.StartTime = utils.TimestampProto(tr.StartTime)
		}
		if tr.EndTime != nil {
			timeRange.EndTime = utils.TimestampProto(tr.EndTime)
		}
		if tr.Step != 0 {
			timeRange.Step = &duration.Duration{Seconds: int64(tr.Step)}
		}
		return &timeRange
	}
	return nil
}

func NewFunction(function *Function) *common.Function {
	if function != nil {
		f := common.Function{}
		f.Type = enumconv.QueryConditionFunctionNameMap[function.Type]
		f.Target = function.Target
		f.Unit = function.Unit
		f.Number = function.Number
		return &f
	}
	return nil
}
