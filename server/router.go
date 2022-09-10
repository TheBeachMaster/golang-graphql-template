package server

import (
	"com.thebeachmaster/golang-graphql-template/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func (srv *Server) MapGraphQLHandlers(app *fiber.App) error {
	gqlRoute := app.Group("")

	handler := routes.NewGraphQLRouteHandler( /*dbClient*/ srv.redisClient, srv.cfg)
	routes.MapGraphQLRoutes(gqlRoute, handler, srv.redisClient)
	return nil
}
