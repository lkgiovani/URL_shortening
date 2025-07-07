package main

import (
	"fmt"
	"log"

	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/delivery/httpserver"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	config, err := environment.NewConfig()
	if err != nil {
		panic(fmt.Errorf("error new config: %w", err))
	}

	db, err := postgres.NewPostgres(config)
	if err != nil {
		panic(fmt.Errorf("error new postgres: %w", err))
	}

	redis, err := redis.NewRedis(config)
	if err != nil {
		panic(fmt.Errorf("error new redis: %w", err))
	}

	app := fiber.New()

	server, err := httpserver.NewServer(app, db, redis, config)
	if err != nil {
		panic(fmt.Errorf("error new server: %w", err))
	}

	server.Router()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", config.HTTP.Url, config.HTTP.Port)))
}
