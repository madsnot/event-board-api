package database

import "github.com/redis/go-redis/v9"

type Client struct {
	DSN    string
	Client *redis.Client
}

func NewNoSqlDbClient(dsn string) Client {
	return Client{
		DSN: dsn,
	}
}

func (c *Client) Open() error {
	opt, err := redis.ParseURL(c.DSN)
	if err != nil {
		return err
	}

	c.Client = redis.NewClient(opt)

	return nil
}
