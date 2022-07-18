package registry

import (
	"com.thebeachmaster/golang-graphql-template/config"
	"com.thebeachmaster/golang-graphql-template/internal/delivery"
	"github.com/go-redis/redis/v8"
)

type registry struct {
	// db *db.Client
	redis *redis.Client
	cfg   *config.Config
}

type Registry interface {
	NewDelivery() delivery.Delivery
}

func New( /*db *db.Client,*/ redis *redis.Client, cfg *config.Config) Registry {
	return &registry{
		// db: dbclient,
		redis: redis,
		cfg:   cfg,
	}
}

func (r *registry) NewDelivery() delivery.Delivery {
	return delivery.Delivery{
		Songs: r.NewSongDelivery(),
	}
}
