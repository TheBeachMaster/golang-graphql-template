package routes

import (
	// "context"
	"net/http"

	"com.thebeachmaster/golang-graphql-template/internal/middlewares"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type GraphQLRouteHandler interface {
	Playground() http.HandlerFunc
	Service() *handler.Server
}

func MapGraphQLRoutes(r fiber.Router, h GraphQLRouteHandler, redis *redis.Client) {
	mw := middlewares.NewMiddlewareManager(redis)

	r.Get("/playground", mw.BasicAuthMW(), func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(h.Playground())(c.Context())
		return nil
	})
	r.Post("/service", func(c *fiber.Ctx) error {
		// ctxKey := middlewares.TokenCtxtKey
		// ctxValue := c.Get("Authorization")
		// ctx := context.WithValue(c.Context(), ctxKey, ctxValue)
		// c.SetUserContext(ctx)
		fasthttpadaptor.NewFastHTTPHandler(h.Service())(c.Context())
		return nil
	})
}
