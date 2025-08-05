package main

import (
	"github.com/charmingruby/drl/config"
	"github.com/charmingruby/drl/pkg/logger"
	"github.com/charmingruby/drl/pkg/redis"
	"github.com/joho/godotenv"
)

func main() {
	log := logger.New()

	if err := godotenv.Load(); err != nil {
		log.Warn("failed to find .env file", "error", err)
	}

	log.Info("loading environment variables...")

	cfg, err := config.New()
	if err != nil {
		log.Error("failed to loading environment variables", "error", err)
		return
	}

	log.Info("environment variables loaded")

	log.Info("connecting to Redis...")

	redis.New(cfg.RedisURI)

	log.Info("connected to Redis")
}
