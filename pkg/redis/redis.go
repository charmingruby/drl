package redis

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Connection *redis.Client
}

func New(addr string) Client {
	cl := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return Client{
		Connection: cl,
	}
}
