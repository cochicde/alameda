package metric

import (
	"fmt"
	EntityPromthNodeCpu "github.com/containers-ai/alameda/datapipe/pkg/entities/prometheus/nodeCPUUsagePercentage"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalPromth "github.com/containers-ai/alameda/internal/pkg/database/prometheus"
	"github.com/pkg/errors"
	"time"
)

// NodeCPUUsagePercentageRepository Repository to access metric node:node_cpu_utilisation:avg1m from prometheus
type NodeCPUUsagePercentageRepository struct {
	PrometheusConfig InternalPromth.Config
}

// NewNodeCPUUsagePercentageRepositoryWithConfig New node cpu usage percentage repository with prometheus configuration
func NewNodeCPUUsagePercentageRepositoryWithConfig(cfg InternalPromth.Config) NodeCPUUsagePercentageRepository {
	return NodeCPUUsagePercentageRepository{PrometheusConfig: cfg}
}

// ListMetricsByPodNamespacedName Provide metrics from response of querying request contain namespace, pod_name and default labels
func (n NodeCPUUsagePercentageRepository) ListMetricsByNodeName(nodeName string, options ...DBCommon.Option) ([]InternalPromth.Entity, error) {

	var (
		err error

		prometheusClient *InternalPromth.Prometheus

		metricName        string
		queryLabelsString string
		queryExpression   string

		response InternalPromth.Response

		entities []InternalPromth.Entity
	)

	prometheusClient, err = InternalPromth.NewClient(&n.PrometheusConfig)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	opt := DBCommon.NewDefaultOptions()
	for _, option := range options {
		option(&opt)
	}

	metricName = EntityPromthNodeCpu.MetricName
	queryLabelsString = n.buildQueryLabelsStringByNodeName(nodeName)

	if queryLabelsString != "" {
		queryExpression = fmt.Sprintf("%s{%s}", metricName, queryLabelsString)
	} else {
		queryExpression = fmt.Sprintf("%s", metricName)
	}

	stepTimeInSeconds := int64(opt.StepTime.Nanoseconds() / int64(time.Second))
	queryExpression, err = InternalPromth.WrapQueryExpression(queryExpression, opt.AggregateOverTimeFunc, stepTimeInSeconds)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	response, err = prometheusClient.QueryRange(queryExpression, opt.StartTime, opt.EndTime, opt.StepTime)
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	} else if response.Status != InternalPromth.StatusSuccess {
		return entities, errors.Errorf("list node cpu usage metrics by node name failed: receive error response from prometheus: %s", response.Error)
	}

	entities, err = response.GetEntities()
	if err != nil {
		return entities, errors.Wrap(err, "list node cpu usage metrics by node name failed")
	}

	return entities, nil
}

func (n NodeCPUUsagePercentageRepository) buildQueryLabelsStringByNodeName(nodeName string) string {

	var (
		queryLabelsString = ""
	)

	if nodeName != "" {
		queryLabelsString += fmt.Sprintf(`%s = "%s"`, EntityPromthNodeCpu.NodeLabel, nodeName)
	}

	return queryLabelsString
}