package middlewares

import "github.com/go-redis/redis/v8"

type Middlewares struct {
	redis *redis.Client
}

func NewMiddlewareManager(redis *redis.Client) *Middlewares {
	return &Middlewares{redis: redis}
}
