package drivers

import (
	"fmt"
	"runtime"

	"com.thebeachmaster/golang-graphql-template/config"
	"github.com/go-redis/redis/v8"
)

func NewRedisDBClient(cfg *config.Config) (*redis.Client, error) {
	host := cfg.Cache.RedisURL
	if host == "" {
		return nil, fmt.Errorf("Redis Host not found")
	}

	redisOpt, err := redis.ParseURL(host)
	if err != nil {
		return nil, err
	}

	redisOpt.PoolSize = runtime.NumCPU()

	redisClient := redis.NewClient(redisOpt)

	return redisClient, nil
}
