package dispatcher

import (
	"context"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/ai-dispatcher/pkg/metrics"
	"github.com/containers-ai/alameda/ai-dispatcher/pkg/queue"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	datahub_common "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common"
	datahub_metrics "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/metrics"
	datahub_predictions "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/predictions"
	datahub_resources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type controllerModelJobSender struct {
	datahubGrpcCn  *grpc.ClientConn
	modelMapper    *ModelMapper
	metricExporter *metrics.Exporter
}

func NewControllerModelJobSender(datahubGrpcCn *grpc.ClientConn, modelMapper *ModelMapper,
	metricExporter *metrics.Exporter) *controllerModelJobSender {
	return &controllerModelJobSender{
		datahubGrpcCn:  datahubGrpcCn,
		modelMapper:    modelMapper,
		metricExporter: metricExporter,
	}
}

func (sender *controllerModelJobSender) sendModelJobs(controllers []*datahub_resources.Controller,
	queueSender queue.QueueSender, pdUnit string, granularity int64, predictionStep int64) {

	datahubServiceClnt := datahub_v1alpha1.NewDatahubServiceClient(sender.datahubGrpcCn)
	for _, controller := range controllers {
		if granularity == 30 && (!viper.GetBool("hourlyPredict") &&
			controller.GetAlamedaControllerSpec().GetScalingTool() !=
				datahub_resources.ScalingTool_VPA) {
			continue
		}

		controllerNS := controller.GetObjectMeta().GetNamespace()
		controllerName := controller.GetObjectMeta().GetName()

		lastPredictionMetrics, err := sender.getLastPrediction(datahubServiceClnt, controller, granularity)
		if err != nil {
			scope.Infof("Get controller %s/%s last prediction failed: %s",
				controllerNS, controllerName, err.Error())
			continue
		}
		if lastPredictionMetrics == nil && err == nil {
			scope.Infof("No prediction found of controller %s/%s",
				controllerNS, controllerName)
		}
		sender.sendJobByMetrics(controller, queueSender, pdUnit, granularity, predictionStep,
			datahubServiceClnt, lastPredictionMetrics)
	}
}

func (sender *controllerModelJobSender) sendJob(controller *datahub_resources.Controller, queueSender queue.QueueSender, pdUnit string,
	granularity int64, controllerInfo *modelInfo) {
	marshaler := jsonpb.Marshaler{}
	dataGranularity := queue.GetGranularityStr(granularity)
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()
	controllerStr, err := marshaler.MarshalToString(controller)
	if err != nil {
		scope.Errorf("Encode pb message failed for %s/%s with granularity seconds %v. %s",
			controllerNS, controllerName, granularity, err.Error())
		return
	}
	if len(controllerInfo.ModelMetrics) > 0 && controllerStr != "" {
		jb := queue.NewJobBuilder(pdUnit, granularity, controllerStr)
		jobJSONStr, err := jb.GetJobJSONString()
		if err != nil {
			scope.Errorf(
				"Prepare model job payload failed for %s/%s with granularity seconds %v. %s",
				controllerNS, controllerName, granularity, err.Error())
			return
		}

		err = queueSender.SendJsonString(modelQueueName, jobJSONStr,
			fmt.Sprintf("%s/%s/%v", controllerNS, controllerName, granularity))
		if err == nil {
			sender.modelMapper.AddModelInfo(pdUnit, dataGranularity, controllerInfo)
		} else {
			scope.Errorf(
				"Send model job payload failed for %s/%s with granularity seconds %v. %s",
				controllerNS, controllerName, granularity, err.Error())
		}
	}
}

func (sender *controllerModelJobSender) genControllerInfo(controllerNS,
	controllerName string, modelMetrics ...datahub_common.MetricType) *modelInfo {
	controllerInfo := new(modelInfo)
	controllerInfo.NamespacedName = &namespacedName{
		Namespace: controllerNS,
		Name:      controllerName,
	}
	controllerInfo.ModelMetrics = modelMetrics
	controllerInfo.SetTimeStamp(time.Now().Unix())
	return controllerInfo
}

func (sender *controllerModelJobSender) getLastPrediction(datahubServiceClnt datahub_v1alpha1.DatahubServiceClient,
	controller *datahub_resources.Controller, granularity int64) ([]*datahub_predictions.MetricData, error) {
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetName()
	controllerPredictRes, err := datahubServiceClnt.ListControllerPredictions(context.Background(),
		&datahub_predictions.ListControllerPredictionsRequest{
			ObjectMeta: []*datahub_resources.ObjectMeta{
				&datahub_resources.ObjectMeta{
					Namespace: controllerNS,
					Name:      controllerName,
				},
			},
			Granularity: granularity,
			QueryCondition: &datahub_common.QueryCondition{
				Limit: 1,
				Order: datahub_common.QueryCondition_DESC,
				TimeRange: &datahub_common.TimeRange{
					Step: &duration.Duration{
						Seconds: granularity,
					},
				},
			},
		})
	if err != nil {
		return nil, err
	}

	if len(controllerPredictRes.GetControllerPredictions()) > 0 {
		lastControllerPrediction := controllerPredictRes.GetControllerPredictions()[0]
		if lastControllerPrediction.GetPredictedRawData() != nil {
			return lastControllerPrediction.GetPredictedRawData(), nil
		} else if lastControllerPrediction.GetPredictedLowerboundData() != nil {
			return lastControllerPrediction.GetPredictedLowerboundData(), nil
		} else if lastControllerPrediction.GetPredictedUpperboundData() != nil {
			return lastControllerPrediction.GetPredictedUpperboundData(), nil
		}
	}
	return nil, nil
}

func (sender *controllerModelJobSender) getQueryMetricStartTime(
	descControllerPredictions []*datahub_predictions.ControllerPrediction) int64 {
	if len(descControllerPredictions) > 0 {
		pdMDs := descControllerPredictions[len(descControllerPredictions)-1].GetPredictedRawData()
		for _, pdMD := range pdMDs {
			mD := pdMD.GetData()
			if len(mD) > 0 {
				return mD[len(mD)-1].GetTime().GetSeconds()
			}
		}
	}
	return 0
}

func (sender *controllerModelJobSender) sendJobByMetrics(controller *datahub_resources.Controller, queueSender queue.QueueSender,
	pdUnit string, granularity int64, predictionStep int64, datahubServiceClnt datahub_v1alpha1.DatahubServiceClient,
	lastPredictionMetrics []*datahub_predictions.MetricData) {
	dataGranularity := queue.GetGranularityStr(granularity)
	queryCondition := &datahub_common.QueryCondition{
		Order: datahub_common.QueryCondition_DESC,
		TimeRange: &datahub_common.TimeRange{
			Step: &duration.Duration{
				Seconds: granularity,
			},
		},
	}
	controllerNS := controller.GetObjectMeta().GetNamespace()
	controllerName := controller.GetObjectMeta().GetNamespace()
	nowSeconds := time.Now().Unix()

	if len(lastPredictionMetrics) == 0 {
		controllerInfo := sender.genControllerInfo(controllerNS, controllerName,
			datahub_common.MetricType_MEMORY_USAGE_BYTES,
			datahub_common.MetricType_DUTY_CYCLE,
		)
		sender.sendJob(controller, queueSender, pdUnit, granularity, controllerInfo)
		scope.Infof("No prediction metric found of controller %s/%s",
			controllerNS, controllerName)
		return
	}
	for _, lastPredictionMetric := range lastPredictionMetrics {
		if len(lastPredictionMetric.GetData()) == 0 {
			controllerInfo := sender.genControllerInfo(controllerNS, controllerName,
				datahub_common.MetricType_MEMORY_USAGE_BYTES,
				datahub_common.MetricType_DUTY_CYCLE)
			sender.sendJob(controller, queueSender, pdUnit, granularity, controllerInfo)
			scope.Infof("No prediction metric %s found of controller %s/%s",
				lastPredictionMetric.GetMetricType().String(), controllerNS, controllerName)
			return
		} else {
			lastPrediction := lastPredictionMetric.GetData()[0]
			lastPredictionTime := lastPredictionMetric.GetData()[0].GetTime().GetSeconds()
			if lastPrediction != nil && lastPredictionTime <= nowSeconds {
				scope.Infof("controller prediction %s/%s is out of date due to last predict time is %v (current: %v)",
					controllerNS, controllerName, lastPredictionTime, nowSeconds)
			}

			if lastPrediction != nil && lastPredictionTime <= nowSeconds {
				controllerInfo := sender.genControllerInfo(controllerNS, controllerName,
					datahub_common.MetricType_MEMORY_USAGE_BYTES,
					datahub_common.MetricType_DUTY_CYCLE)
				scope.Infof("send controller %s/%s model job due to no predict found or predict is out of date",
					controllerNS, controllerName)
				sender.sendJob(controller, queueSender, pdUnit, granularity, controllerInfo)
				return
			}
			controllerPredictRes, err := datahubServiceClnt.ListControllerPredictions(context.Background(),
				&datahub_predictions.ListControllerPredictionsRequest{
					ObjectMeta: []*datahub_resources.ObjectMeta{
						&datahub_resources.ObjectMeta{
							Namespace: controllerNS,
							Name:      controllerName,
						},
					},
					Granularity:    granularity,
					ModelId:        lastPrediction.GetModelId(),
					QueryCondition: queryCondition,
				})
			if err != nil {
				scope.Errorf("Get controller %s/%s Prediction with granularity %v for sending model job failed: %s",
					controllerNS, controllerName, granularity, err.Error())
				continue
			}
			controllerPredictions := controllerPredictRes.GetControllerPredictions()
			queryStartTime := time.Now().Unix() - predictionStep*granularity
			firstPDTime := sender.getQueryMetricStartTime(controllerPredictions)
			if firstPDTime > 0 {
				queryStartTime = firstPDTime
			}
			controllerMetricsRes, err := datahubServiceClnt.ListControllerMetrics(context.Background(),
				&datahub_metrics.ListControllerMetricsRequest{
					QueryCondition: &datahub_common.QueryCondition{
						Order: datahub_common.QueryCondition_DESC,
						TimeRange: &datahub_common.TimeRange{
							StartTime: &timestamp.Timestamp{
								Seconds: queryStartTime,
							},
							Step: &duration.Duration{
								Seconds: granularity,
							},
						},
					},
					ObjectMeta: []*datahub_resources.ObjectMeta{
						&datahub_resources.ObjectMeta{
							Namespace: controllerNS,
							Name:      controllerName,
						},
					},
				})

			if err != nil {
				scope.Errorf("List controller %s/%s metric with granularity %v for sending model job failed: %s",
					controllerNS, controllerName, granularity, err.Error())
				continue
			}
			controllerMetrics := controllerMetricsRes.GetControllerMetrics()

			for _, controllerMetric := range controllerMetrics {
				metricData := controllerMetric.GetMetricData()
				for _, metricDatum := range metricData {
					mData := metricDatum.GetData()
					pData := []*datahub_predictions.Sample{}
					controllerInfo := sender.genControllerInfo(controllerNS, controllerName)
					for _, controllerPrediction := range controllerPredictions {
						predictRawData := controllerPrediction.GetPredictedRawData()
						for _, predictRawDatum := range predictRawData {
							if metricDatum.GetMetricType() == predictRawDatum.GetMetricType() {
								pData = append(pData, predictRawDatum.GetData()...)
							}
						}
					}
					metricsNeedToModel, drift := DriftEvaluation(UnitTypeController, metricDatum.GetMetricType(), granularity, mData, pData, map[string]string{
						"controllerNS":   controllerNS,
						"controllerName": controllerName,
						"controllerKind": controller.GetKind().String(),
						"targetDisplayName": fmt.Sprintf("controller %s %s/%s", controller.GetKind().String(),
							controllerNS, controllerName),
					}, sender.metricExporter)
					if drift {
						scope.Infof("export controller %s %s/%s drift counter with granularity %s",
							controller.GetKind().String(), controllerNS, controllerName, dataGranularity)
						sender.metricExporter.AddControllerMetricDrift(controllerNS, controllerName,
							controller.GetKind().String(), queue.GetGranularityStr(granularity), 1.0)
					}
					controllerInfo.ModelMetrics = append(controllerInfo.ModelMetrics, metricsNeedToModel...)
					isModeling := sender.modelMapper.IsModeling(pdUnit, dataGranularity, controllerInfo)
					if !isModeling || (isModeling && sender.modelMapper.IsModelTimeout(
						pdUnit, dataGranularity, controllerInfo)) {
						sender.sendJob(controller, queueSender, pdUnit, granularity, controllerInfo)
						return
					}
				}
			}
		}
	}
}