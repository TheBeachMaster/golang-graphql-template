package routes

import (
	"net/http"

	"com.thebeachmaster/golang-graphql-template/config"
	"com.thebeachmaster/golang-graphql-template/pkg/graphql"
	"com.thebeachmaster/golang-graphql-template/pkg/registry"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis/v8"
)

type graphqlRouteHandler struct {
	// db   *db.Client
	redis *redis.Client
	cfg   *config.Config
}

func NewGraphQLRouteHandler( /*DB Client*/ redis *redis.Client, cfg *config.Config) GraphQLRouteHandler {
	return &graphqlRouteHandler{ /*db: client,*/ redis: redis, cfg: cfg}
}

func (r *graphqlRouteHandler) Playground() http.HandlerFunc {
	pg := playground.Handler("TawiGraphQL Playground", "/service")
	return pg
}

func (r *graphqlRouteHandler) Service() *handler.Server {
	serviceRegistry := registry.New( /*db client*/ r.redis, r.cfg)
	registryController := serviceRegistry.NewDelivery()
	srv := graphql.NewGraphQLServer( /*db client*/ r.redis, &registryController, r.cfg)

	return srv
}
