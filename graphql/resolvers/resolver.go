package resolvers

import (
	"com.thebeachmaster/golang-graphql-template/config"
	"com.thebeachmaster/golang-graphql-template/graphql/generated"
	"com.thebeachmaster/golang-graphql-template/internal/delivery"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-redis/redis/v8"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	// db *db.Client
	redis    *redis.Client
	delivery delivery.Delivery
	config   *config.Config
}

func NewSchemaReslover( /*client *ent.Client,*/ redis *redis.Client, dlvr delivery.Delivery, cfg *config.Config) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			// db:   db.client,
			delivery: dlvr,
			redis:    redis,
			config:   cfg,
		},
	})
}
