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
)

type CreateClusterPredictionsRequestExtended struct {
	*ApiPredictions.CreateClusterPredictionsRequest
}

func (r *CreateClusterPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *CreateClusterPredictionsRequestExtended) ProducePredictions() DaoPredictionTypes.ClusterPredictionMap {
	clusterPredictionMap := DaoPredictionTypes.NewClusterPredictionMap()

	for _, cluster := range r.GetClusterPredictions() {
		// Normalize request
		objectMeta := resources.NewObjectMeta(cluster.GetObjectMeta())
		objectMeta.Namespace = ""
		objectMeta.NodeName = ""
		objectMeta.ClusterName = ""
		objectMeta.NodeName = ""

		clusterPrediction := DaoPredictionTypes.NewClusterPrediction()
		clusterPrediction.ObjectMeta.Name = cluster.GetObjectMeta().GetName()

		// Handle predicted raw data
		for _, data := range cluster.GetPredictedRawData() {
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
				clusterPrediction.AddRawSample(metricType, granularity, sample)
			}
		}

		// Handle predicted upper bound data
		for _, data := range cluster.GetPredictedUpperboundData() {
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
				clusterPrediction.AddUpperBoundSample(metricType, granularity, sample)
			}
		}

		// Handle predicted lower bound data
		for _, data := range cluster.GetPredictedLowerboundData() {
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
				clusterPrediction.AddLowerBoundSample(metricType, granularity, sample)
			}
		}

		clusterPredictionMap.AddClusterPrediction(clusterPrediction)
	}

	return clusterPredictionMap
}

type ListClusterPredictionsRequestExtended struct {
	Request *ApiPredictions.ListClusterPredictionsRequest
}

func (r *ListClusterPredictionsRequestExtended) Validate() error {
	return nil
}

func (r *ListClusterPredictionsRequestExtended) ProduceRequest() DaoPredictionTypes.ListClusterPredictionsRequest {
	request := DaoPredictionTypes.NewListClusterPredictionRequest()
	request.QueryCondition = common.QueryConditionExtend{Condition: r.Request.GetQueryCondition()}.QueryCondition()
	request.Granularity = 30
	request.ModelId = r.Request.GetModelId()
	request.PredictionId = r.Request.GetPredictionId()
	if r.Request.GetGranularity() != 0 {
		request.Granularity = r.Request.GetGranularity()
	}
	if r.Request.GetObjectMeta() != nil {
		for _, meta := range r.Request.GetObjectMeta() {
			// Normalize request
			objectMeta := resources.NewObjectMeta(meta)
			objectMeta.Namespace = ""
			objectMeta.NodeName = ""
			objectMeta.ClusterName = ""

			if objectMeta.IsEmpty() {
				request.ObjectMeta = make([]Metadata.ObjectMeta, 0)
				return request
			}
			request.ObjectMeta = append(request.ObjectMeta, objectMeta)
		}
	}
	return request
}
