package core_redis_pool

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"ADDR" default:"localhost:6379"`
	Password string `envconfig:"PASSWORD" required:"true"`
	DB int `envconfig:"DB" default:"0"`
	Timeout time.Duration `envconfig:"TIMEOUT" default:"5s"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("REDIS", &config); err != nil {
		return Config{}, fmt.Errorf("proccess envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err := fmt.Errorf("get redis client config: %w", err)
		panic(err)
	}
	return config
}