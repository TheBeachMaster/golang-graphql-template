package server

import (
	"log"
	"os"
	"os/signal"
	"time"

	"com.thebeachmaster/golang-graphql-template/config"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app         *fiber.App
	cfg         *config.Config
	redisClient *redis.Client
	//dbClient
}

func NewServer(cfg *config.Config /*dbClient*/, redis *redis.Client) *Server {

	_app := fiber.New(fiber.Config{
		Prefork:      cfg.Server.Prefork,
		ReadTimeout:  time.Second * time.Duration(cfg.Server.ReadTimeout),
		AppName:      cfg.Server.AppName,
		ServerHeader: cfg.Server.ServerHeader,
	})

	_app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	return &Server{app: _app, cfg: cfg, redisClient: redis}
}

func (srv *Server) Run() error {

	go func() {
		log.Printf("Server is listening on PORT: %s\n", srv.cfg.Port)
		addr := ":" + srv.cfg.Port
		if err := srv.app.Listen(addr); err != nil {
			log.Panicf("[CRIT] Unable to start server. Reason: %v", err)
		}
	}()

	err := srv.MapGraphQLHandlers(srv.app)
	if err != nil {
		return err
	}

	quitServer := make(chan struct{})

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	close(quitServer)

	<-quitServer

	log.Printf("Server shutdown")
	return srv.app.Shutdown()

}
