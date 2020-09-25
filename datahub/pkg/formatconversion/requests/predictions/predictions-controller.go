package predictions

import (
	"github.com/golang/protobuf/ptypes"
	DaoPredictionTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/predictions/types"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/common"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/metrics"
	"prophetstor.com/alameda/datahub/pkg/formatconversion/requests/resources"
	FormatTypes "prophetstor.com/alameda/datahub/pkg/formatconversion/types"
	Metadata "prophetstor.com/alameda/datahub/pkg/kubernetes/metadata"
	ApiPredictions "prophetstor.com/api/datahub/predictions"
	ApiResources "prophetstor.com/api/datahub/resources"
)

type CreateControllerPredictionsRequestExtended struct {
	*ApiPredictions.CreateControllerPredictionsRequest
}

func (r *CreateControllerPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *CreateControllerPredictionsRequestExtended) ProducePredictions() DaoPredictionTypes.ControllerPredictionMap {
	controllerPredictionMap := DaoPredictionTypes.NewControllerPredictionMap()

	for _, controller := range r.GetControllerPredictions() {
		// Normalize request
		objectMeta := resources.NewObjectMeta(controller.GetObjectMeta())
		objectMeta.NodeName = ""

		controllerPrediction := DaoPredictionTypes.NewControllerPrediction()
		controllerPrediction.ObjectMeta = objectMeta
		controllerPrediction.Kind = controller.GetKind().String()

		// Handle predicted raw data
		for _, data := range controller.GetPredictedRawData() {
			metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
			granularity := data.GetGranularity()

			for _, sample := range data.GetData() {
				timestamp, err := ptypes.Timestamp(sample.GetTime())
				if err != nil {
					scope.Error(" failed: " + err.Error())
				}
				sample := FormatTypes.PredictionSample{
					Timestamp:    timestamp,
					Value:        sample.GetNumValue(),
					ModelId:      sample.GetModelId(),
					PredictionId: sample.GetPredictionId(),
				}
				controllerPrediction.AddRawSample(metricType, granularity, sample)
			}
		}

		// Handle predicted upper bound data
		for _, data := range controller.GetPredictedUpperboundData() {
			metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
			granularity := data.GetGranularity()
			for _, sample := range data.GetData() {
				timestamp, err := ptypes.Timestamp(sample.GetTime())
				if err != nil {
					scope.Error(" failed: " + err.Error())
				}
				sample := FormatTypes.PredictionSample{
					Timestamp:    timestamp,
					Value:        sample.GetNumValue(),
					ModelId:      sample.GetModelId(),
					PredictionId: sample.GetPredictionId(),
				}
				controllerPrediction.AddUpperBoundSample(metricType, granularity, sample)
			}
		}

		// Handle predicted lower bound data
		for _, data := range controller.GetPredictedLowerboundData() {
			metricType := metrics.MetricTypeNameMap[data.GetMetricType()]
			granularity := data.GetGranularity()
			for _, sample := range data.GetData() {
				timestamp, err := ptypes.Timestamp(sample.GetTime())
				if err != nil {
					scope.Error(" failed: " + err.Error())
				}
				sample := FormatTypes.PredictionSample{
					Timestamp:    timestamp,
					Value:        sample.GetNumValue(),
					ModelId:      sample.GetModelId(),
					PredictionId: sample.GetPredictionId(),
				}
				controllerPrediction.AddLowerBoundSample(metricType, granularity, sample)
			}
		}

		controllerPredictionMap.AddControllerPrediction(controllerPrediction)
	}

	return controllerPredictionMap
}

type ListControllerPredictionsRequestExtended struct {
	Request *ApiPredictions.ListControllerPredictionsRequest
}

func (r *ListControllerPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *ListControllerPredictionsRequestExtended) ProduceRequest() DaoPredictionTypes.ListControllerPredictionsRequest {
	request := DaoPredictionTypes.NewListControllerPredictionRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.Granularity = 30
	request.ModelId = r.Request.GetModelId()
	request.PredictionId = r.Request.GetPredictionId()
	if r.Request.GetKind() != ApiResources.Kind_KIND_UNDEFINED {
		request.Kind = r.Request.GetKind().String()
	}
	if r.Request.GetGranularity() != 0 {
		request.Granularity = r.Request.GetGranularity()
	}
	if r.Request.GetObjectMeta() != nil {
		for _, meta := range r.Request.GetObjectMeta() {
			// Normalize request
			objectMeta := resources.NewObjectMeta(meta)
			objectMeta.NodeName = ""

			if objectMeta.IsEmpty() {
				request.ObjectMeta = make([]Metadata.ObjectMeta, 0)
				return request
			}
			request.ObjectMeta = append(request.ObjectMeta, objectMeta)
		}
	}
	return request
}
