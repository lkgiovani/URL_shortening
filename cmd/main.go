package main

import (
	"fmt"
	"log"

	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/internal/delivery/http"

	"github.com/gofiber/fiber/v3"
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

	app := fiber.New()

	server, err := http.NewServer(app, db)
	if err != nil {
		panic(fmt.Errorf("error new server: %w", err))
	}

	server.Router()

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", config.HTTP.Url, config.HTTP.Port)))
}
