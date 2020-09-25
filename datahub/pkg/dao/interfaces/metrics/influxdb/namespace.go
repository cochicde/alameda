package influxdb

import (
	"context"

	"github.com/pkg/errors"
	DaoMetricTypes "prophetstor.com/alameda/datahub/pkg/dao/interfaces/metrics/types"
	RepoInfluxMetric "prophetstor.com/alameda/datahub/pkg/dao/repositories/influxdb/metrics"
	FormatEnum "prophetstor.com/alameda/datahub/pkg/formatconversion/enumconv"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	Utils "prophetstor.com/alameda/pkg/utils"
)

type NamespaceMetrics struct {
	InfluxDBConfig InfluxDB.Config
}

func NewNamespaceMetricsWithConfig(config InfluxDB.Config) DaoMetricTypes.NamespaceMetricsDAO {
	return &NamespaceMetrics{InfluxDBConfig: config}
}

func (n NamespaceMetrics) CreateMetrics(ctx context.Context, m DaoMetricTypes.NamespaceMetricMap) error {
	// Write namespace cpu metrics
	cpuRepo := RepoInfluxMetric.NewNamespaceCPURepositoryWithConfig(n.InfluxDBConfig)
	err := cpuRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeCPUUsageSecondsPercentage))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create namespace cpu metrics failed")
	}

	// Write namespace memory metrics
	memoryRepo := RepoInfluxMetric.NewNamespaceMemoryRepositoryWithConfig(n.InfluxDBConfig)
	err = memoryRepo.CreateMetrics(ctx, m.GetSamples(FormatEnum.MetricTypeMemoryUsageBytes))
	if err != nil {
		scope.Error(err.Error())
		return errors.Wrap(err, "create namespace memory metrics failed")
	}
	return nil
}

func (n NamespaceMetrics) ListMetrics(ctx context.Context, req DaoMetricTypes.ListNamespaceMetricsRequest) (DaoMetricTypes.NamespaceMetricMap, error) {
	metricMap := DaoMetricTypes.NewNamespaceMetricMap()

	// Read namespace cpu metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeCPUUsageSecondsPercentage) {
		cpuRepo := RepoInfluxMetric.NewNamespaceCPURepositoryWithConfig(n.InfluxDBConfig)
		cpuMetricMap, err := cpuRepo.GetNamespaceMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get namespace cpu usage metric map failed")
		}
		for _, m := range cpuMetricMap.MetricMap {
			copyM := m
			metricMap.AddNamespaceMetric(copyM)
		}
	}

	// Read namespace memory metrics
	if Utils.SliceContains(req.MetricTypes, FormatEnum.MetricTypeMemoryUsageBytes) {
		memoryRepo := RepoInfluxMetric.NewNamespaceMemoryRepositoryWithConfig(n.InfluxDBConfig)
		memoryMetricMap, err := memoryRepo.GetNamespaceMetricMap(ctx, req)
		if err != nil {
			scope.Error(err.Error())
			return metricMap, errors.Wrap(err, "get namespace memory usage metric map failed")
		}
		for _, m := range memoryMetricMap.MetricMap {
			copyM := m
			metricMap.AddNamespaceMetric(copyM)
		}
	}

	return metricMap, nil

}
