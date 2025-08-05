package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/charmingruby/drl/example/config"
	"github.com/charmingruby/drl/example/http"
	"github.com/charmingruby/drl/pkg/logger"
	"github.com/charmingruby/drl/pkg/rate_limiter"
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

	redisCl := redis.New(cfg.RedisURI)

	log.Info("connected to Redis")

	rateLimiter := rate_limiter.New(&redisCl, 5, 60)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)

	srv := http.NewServer(addr)

	mw := http.NewMiddleware(&rateLimiter, log)

	http.RegisterRoutes(srv.Router, mw)

	log.Info("starting server...")

	go func() {
		log.Info("server is running", "port", cfg.ServerPort)

		if err := srv.Start(); err != nil {
			log.Error("server error", "error", err.Error())
			return
		}
	}()

	exit := make(chan os.Signal, 1)

	signal.Notify(exit, os.Interrupt)

	signal := <-exit

	log.Info("received signal", "signal", signal)

	log.Info("application shutdown")
}
