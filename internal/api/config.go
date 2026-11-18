package api

import (
	"uuid"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName     string `default:"bookmark-service" envconfig:"APP_NAME"`
	AppPort     int    `default:"8081" envconfig:"APP_PORT"`
	ServiceName string `default:"bookmark_service" envconfig:"SERVICE_NAME"`
	InstanceId  string `envconfig:"INSTANCE_ID"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	if cfg.InstanceId == "" {
		cfg.InstanceId = uuid.New().String()
	}
	return cfg, nil
}
