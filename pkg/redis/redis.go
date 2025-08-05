package redis

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Conn *redis.Client
}

func New(addr string) Client {
	cl := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return Client{
		Conn: cl,
	}
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
