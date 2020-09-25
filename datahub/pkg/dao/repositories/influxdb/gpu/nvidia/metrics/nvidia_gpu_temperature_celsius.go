package metrics

import (
	"fmt"
	Client "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
	EntityInfluxGpuMetric "prophetstor.com/alameda/datahub/pkg/dao/entities/influxdb/gpu/nvidia/metrics"
	RepoInflux "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	InfluxModels "prophetstor.com/alameda/pkg/database/influxdb/models"
)

type TemperatureCelsiusRepository struct {
	influxDB *InfluxDB.InfluxClient
}

func NewTemperatureCelsiusRepositoryWithConfig(cfg InfluxDB.Config) *TemperatureCelsiusRepository {
	return &TemperatureCelsiusRepository{
		influxDB: InfluxDB.NewClient(&cfg),
	}
}

func (r *TemperatureCelsiusRepository) ListMetrics(host, minorNumber string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, error) {
	steps := int(condition.StepTime.Seconds())
	if steps == 0 || steps == 30 {
		return r.read(host, minorNumber, condition)
	} else {
		return r.steps(host, minorNumber, condition)
	}
}

func (r *TemperatureCelsiusRepository) read(host, minorNumber string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, error) {
	entities := make([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, 0)

	influxdbStatement := InfluxDB.Statement{
		QueryCondition: condition,
		Measurement:    TemperatureCelsius,
		GroupByTags:    []string{"host"},
	}

	influxdbStatement.AppendWhereClause("AND", "host", "=", host)
	influxdbStatement.AppendWhereClause("AND", "minor_number", "=", minorNumber)
	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()
	cmd := influxdbStatement.BuildQueryCmd()

	response, err := r.influxDB.QueryDB(cmd, string(RepoInflux.Gpu))
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu temperature celsius")
	}

	entities = r.genEntities(response)

	return entities, nil
}

func (r *TemperatureCelsiusRepository) steps(host, minorNumber string, condition *DBCommon.QueryCondition) ([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, error) {
	entities := make([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, 0)

	response, err := r.last(host, minorNumber, condition)
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu temperature celsius with last")
	}
	lastEntities := r.genEntities(response)

	response, err = r.max(host, minorNumber, condition)
	if err != nil {
		return entities, errors.Wrap(err, "failed to list nvidia gpu temperature celsius with max")
	}
	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			entityPtr := &EntityInfluxGpuMetric.TemperatureCelsiusEntity{}
			group := result.GetGroup(i)
			gpuId := group.Tags["uuid"]
			found := false

			for _, entityPtr = range lastEntities {
				if *entityPtr.Uuid == gpuId {
					found = true
					break
				}
			}

			if found {
				for j := 0; j < group.GetRowNum(); j++ {
					row := group.GetRow(j)
					if row["max_value"] != "" {
						entityMap := make(map[string]string)
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusTime] = row["time"]
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusHost] = *entityPtr.Host
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusInstance] = *entityPtr.Instance
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusJob] = *entityPtr.Job
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusName] = *entityPtr.Name
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusUuid] = *entityPtr.Uuid

						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusMinorNumber] = *entityPtr.MinorNumber
						entityMap[EntityInfluxGpuMetric.TemperatureCelsiusValue] = row["max_value"]

						entity := EntityInfluxGpuMetric.NewTemperatureCelsiusEntityFromMap(entityMap)
						entities = append(entities, &entity)
					}
				}
			}
		}
	}

	return entities, nil
}

func (r *TemperatureCelsiusRepository) last(host, minorNumber string, condition *DBCommon.QueryCondition) ([]Client.Result, error) {
	queryCondition := *condition
	queryCondition.Limit = 1

	statement := InfluxDB.Statement{
		QueryCondition: &queryCondition,
		Measurement:    TemperatureCelsius,
		GroupByTags:    []string{"uuid"},
	}

	statement.AppendWhereClause("AND", "host", "=", host)
	statement.AppendWhereClause("AND", "minor_number", "=", minorNumber)
	statement.AppendWhereClauseFromTimeCondition()
	statement.SetOrderClauseFromQueryCondition()
	statement.SetLimitClauseFromQueryCondition()
	cmd := statement.BuildQueryCmd()

	return r.influxDB.QueryDB(cmd, string(RepoInflux.Gpu))
}

func (r *TemperatureCelsiusRepository) max(host, minorNumber string, condition *DBCommon.QueryCondition) ([]Client.Result, error) {
	seconds := int(condition.StepTime.Seconds())
	groupTag := fmt.Sprintf("time(%ds)", seconds)

	statement := InfluxDB.Statement{
		QueryCondition: condition,
		Measurement:    TemperatureCelsius,
		GroupByTags:    []string{"uuid", groupTag},
	}

	statement.AppendWhereClause("AND", "host", "=", host)
	statement.AppendWhereClause("AND", "minor_number", "=", minorNumber)
	statement.AppendWhereClauseFromTimeCondition()
	statement.SetOrderClauseFromQueryCondition()
	statement.SetLimitClauseFromQueryCondition()
	statement.SetFunction(InfluxDB.Select, "MAX", "")
	cmd := statement.BuildQueryCmd()

	return r.influxDB.QueryDB(cmd, string(RepoInflux.Gpu))
}

func (r *TemperatureCelsiusRepository) genEntities(response []Client.Result) []*EntityInfluxGpuMetric.TemperatureCelsiusEntity {
	entities := make([]*EntityInfluxGpuMetric.TemperatureCelsiusEntity, 0)

	results := InfluxModels.NewInfluxResults(response)
	for _, result := range results {
		for i := 0; i < result.GetGroupNum(); i++ {
			group := result.GetGroup(i)
			for j := 0; j < group.GetRowNum(); j++ {
				entity := EntityInfluxGpuMetric.NewTemperatureCelsiusEntityFromMap(group.GetRow(j))
				entities = append(entities, &entity)
			}
		}
	}

	return entities
}
