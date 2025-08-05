package config

import "github.com/caarlos0/env"

type Config struct {
	RedisURI string `env:"REDIS_URI,required"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
