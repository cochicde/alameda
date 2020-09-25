package main

import (
	"prophetstor.com/alameda/internal/pkg/message-queue/kafka"
	"prophetstor.com/alameda/operator/datahub"
	"prophetstor.com/alameda/pkg/utils/log"

	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// Config defines configurations
type Config struct {
	Log     *log.Config     `mapstructure:"log"`
	Datahub *datahub.Config `mapstructure:"datahub"`
	Manager manager.Manager

	Kafka *kafka.Config `mapstructure:"kafka"`
}

// NewConfig returns Config objecdt
func NewConfig(manager manager.Manager) Config {

	c := Config{
		Manager: manager,
	}
	c.init()

	return c
}

func NewConfigWithoutMgr() Config {

	c := Config{}
	c.init()

	return c
}

func (c *Config) init() {

	defaultLogConfig := log.NewDefaultConfig()

	c.Log = &defaultLogConfig
	c.Datahub = datahub.NewConfig()
}

func (c Config) Validate() error {
	return nil
}
