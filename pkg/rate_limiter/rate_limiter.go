package rate_limiter

import (
	"context"
	"time"
)

type Storage interface {
	Store()
	Get()
}

type RateLimiter struct {
	ctx     context.Context
	storage Storage
	window  time.Duration
	limit   int
}

func New(st Storage, limit, windowInSeconds int) RateLimiter {
	windowDuration := time.Duration(windowInSeconds) * time.Second

	return RateLimiter{
		ctx:     context.Background(),
		storage: st,
		limit:   limit,
		window:  windowDuration,
	}
}
