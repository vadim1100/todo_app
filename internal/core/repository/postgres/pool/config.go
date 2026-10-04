package core_postgres_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
    PostgresUser     string `envconfig:"USER" required:"true"`
    PostgresPassword string `envconfig:"PASSWORD" required:"true"`
    PostgresDB       string `envconfig:"DB" required:"true"`
    PostgresHost     string `envconfig:"HOST" default:"localhost"`
    PostgresPort     int    `envconfig:"PORT" default:"5432"`
	Timeout time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get postgres pool config: %w", err)
		panic(err)
	}
	return config
}