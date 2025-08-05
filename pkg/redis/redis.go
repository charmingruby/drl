package redis

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	conn *redis.Client
}

func New(addr string) Client {
	cl := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return Client{
		conn: cl,
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}
