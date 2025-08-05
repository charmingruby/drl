package rate_limiter

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/charmingruby/drl/pkg/redis"
)

var (
	ErrRedisExecution = errors.New("failed executing redis command")
)

type RateLimiter struct {
	ctx    context.Context
	redis  *redis.Client
	window time.Duration
	limit  int
}

func New(redis *redis.Client, limit, windowInSeconds int) RateLimiter {
	windowDuration := time.Duration(windowInSeconds) * time.Second

	return RateLimiter{
		ctx:    context.Background(),
		redis:  redis,
		limit:  limit,
		window: windowDuration,
	}
}

func (rt *RateLimiter) Allow(key string) (bool, error) {
	pipe := rt.redis.Conn.TxPipeline()

	incr := pipe.Incr(rt.ctx, key)

	pipe.Expire(rt.ctx, key, rt.window)

	_, err := pipe.Exec(rt.ctx)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrRedisExecution, err)
	}

	return incr.Val() <= int64(rt.limit), nil
}
