package core_kafka

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Brokers []string `envconfig:"BROKERS" required:"true"`
	Topic   string   `envconfig:"TOPIC" required:"true"`
	GroupID string   `envconfig:"GROUP_ID" default:"audit-log"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("KAFKA", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get kafka config: %w", err)
		panic(err)
	}
	return config
}
