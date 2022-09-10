package main

import (
	"log"
	"os"

	"com.thebeachmaster/golang-graphql-template/config"
	"com.thebeachmaster/golang-graphql-template/drivers"
	"com.thebeachmaster/golang-graphql-template/server"
)

func main() {
	log.Println("Starting Server...")

	var cfg *config.Config

	if os.Getenv("SERVER_ENV") == "production" {

		appConfigPath := "./config/config"

		cfgFile, err := config.LoadConfig(appConfigPath)
		if err != nil {
			log.Fatalf("LoadConfig Error: %v", err)
		}

		pc, err := config.ParseConfig(cfgFile)
		if err != nil {
			log.Fatalf("ParseConfig Error: %v", err)
		}

		cfg = pc

	} else {

		lc, err := config.LoadEnvConfig()
		if err != nil {
			log.Fatalf("ParseConfig Error: %v", err)
		}
		cfg = lc
	}

	redisClient, err := drivers.NewRedisDBClient(cfg)
	if err != nil {
		log.Printf("failed to connect to redis due to: %v\n", err)
	}
	defer redisClient.Close()

	server := server.NewServer(cfg /**dbclient*/, redisClient)
	if err = server.Run(); err != nil {
		log.Fatal(err)
	}
}
