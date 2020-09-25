package probe

import (
	"os"
	"prophetstor.com/alameda/pkg/utils/log"
)

var scope = log.RegisterScope("probe", "datahub health probe", 0)

func LivenessProbe(cfg *LivenessProbeConfig) {
	bindAddr := cfg.BindAddr
	err := queryDatahub(bindAddr)
	if err != nil {
		scope.Errorf("Liveness probe failed with address (%s) due to %s", bindAddr, err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func ReadinessProbe(cfg *ReadinessProbeConfig) {
	influxdbCfg := cfg.InfluxdbCfg
	queueCfg := cfg.RabbitMQCfg

	err := queryInfluxdb(influxdbCfg)
	if err != nil {
		scope.Errorf("Readiness probe: failed to ping influxdb with address (%s) due to %s", influxdbCfg.Address, err.Error())
		os.Exit(1)
	}

	err = queryQueue(queueCfg)
	if err != nil {
		scope.Errorf("Readiness probe: failed to query queue with url (%s) due to %s", queueCfg.URL, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
