package core_service_jwt

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	JWTSecret string        `envconfig:"JWT_SECRET" required:"true"`
    JWTTTL    time.Duration `envconfig:"JWT_TTL" default:"24h"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("JWT", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get jwt config: %w", err)
		panic(err)
	}
	return config
}