package graphql

import (
	config "com.thebeachmaster/golang-graphql-template/config"
	"com.thebeachmaster/golang-graphql-template/graphql/resolvers"
	"com.thebeachmaster/golang-graphql-template/internal/delivery"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-redis/redis/v8"
)

func NewGraphQLServer( /*DB Client*/ redis *redis.Client, delivery *delivery.Delivery, cfg *config.Config) *handler.Server {
	var mb int64 = 1 << 20
	gqlSrv := handler.NewDefaultServer(resolvers.NewSchemaReslover(redis, *delivery, cfg))
	gqlSrv.AddTransport(transport.POST{})
	gqlSrv.AddTransport(transport.MultipartForm{
		MaxMemory:     32 * mb,
		MaxUploadSize: 50 * mb,
	})
	return gqlSrv
}
