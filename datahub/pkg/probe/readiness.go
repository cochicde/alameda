package probe

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"os/exec"
	RepoPromthMetric "prophetstor.com/alameda/datahub/pkg/dao/repositories/prometheus/metrics"
	InternalRabbitMQ "prophetstor.com/alameda/internal/pkg/message-queue/rabbitmq"
	DBCommon "prophetstor.com/alameda/pkg/database/common"
	InfluxDB "prophetstor.com/alameda/pkg/database/influxdb"
	Prometheus "prophetstor.com/alameda/pkg/database/prometheus"
	"strings"
)

type ReadinessProbeConfig struct {
	InfluxdbCfg   *InfluxDB.Config
	PrometheusCfg *Prometheus.Config
	RabbitMQCfg   *InternalRabbitMQ.Config
}

func queryInfluxdb(influxdbConfig *InfluxDB.Config) error {
	err := pingInfluxdb(influxdbConfig.Address)
	if err != nil {
		return errors.Wrap(err, "failed to ping to influxdb")
	}
	return nil
}

func queryPrometheus(prometheusConfig *Prometheus.Config) error {
	emr := "exceeded maximum resolution"
	options := []DBCommon.Option{}
	ctx := context.Background()

	podContainerCPURepo := RepoPromthMetric.NewContainerCpuUsageRepositoryWithConfig(*prometheusConfig)
	containerCPUEntities, err := podContainerCPURepo.ListContainerCPUUsageMillicoresEntitiesByNamespaceAndPodNames(ctx, "", nil, options...)
	if err != nil && !strings.Contains(err.Error(), emr) {
		return errors.Wrap(err, "list pod metrics failed")
	}

	if err == nil && len(containerCPUEntities) == 0 {
		return fmt.Errorf("no container CPU metric found")
	}

	podContainerMemoryRepo := RepoPromthMetric.NewContainerMemoryUsageRepositoryWithConfig(*prometheusConfig)
	containerMemoryEntities, err := podContainerMemoryRepo.ListContainerMemoryUsageBytesEntitiesByNamespaceAndPodNames(ctx, "", nil, options...)
	if err != nil && !strings.Contains(err.Error(), emr) {
		return errors.Wrap(err, "list pod metrics failed")
	}

	if err == nil && len(containerMemoryEntities) == 0 {
		return fmt.Errorf("no container memory metric found")
	}

	nodeCPUUsageRepo := RepoPromthMetric.NewNodeCPUUsageRepositoryWithConfig(*prometheusConfig)
	nodeCPUUsageEntities, err := nodeCPUUsageRepo.ListNodeCPUUsageMillicoresEntitiesByNodeNames(context.TODO(), nil, options...)
	if err != nil && !strings.Contains(err.Error(), emr) {
		return errors.Wrap(err, "list node cpu usage metrics failed")
	}

	if err == nil && len(nodeCPUUsageEntities) == 0 {
		return fmt.Errorf("no node CPU metric found")
	}

	nodeMemoryUsageRepo := RepoPromthMetric.NewNodeMemoryUsageRepositoryWithConfig(*prometheusConfig)
	nodeMemoryUsageEntities, err := nodeMemoryUsageRepo.ListNodeMemoryBytesUsageEntitiesByNodeNames(context.TODO(), nil, options...)
	if err != nil && !strings.Contains(err.Error(), emr) {
		return errors.Wrap(err, "list node memory usage metrics failed")
	}

	if err == nil && len(nodeMemoryUsageEntities) == 0 {
		return fmt.Errorf("no node memory metric found")
	}

	return nil
}

func queryQueue(rabbitmqConfig *InternalRabbitMQ.Config) error {
	return connQueue(rabbitmqConfig.URL)
}

func connQueue(url string) error {
	_, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	return nil
}

func pingInfluxdb(influxdbAddr string) error {
	pingURL := fmt.Sprintf("%s/ping", influxdbAddr)
	curlCmd := exec.Command("curl", "-sl", "-I", pingURL)
	if strings.Contains(pingURL, "https") {
		curlCmd = exec.Command("curl", "-sl", "-I", "-k", pingURL)
	}
	_, err := curlCmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}
