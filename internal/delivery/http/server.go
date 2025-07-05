package http

import (
	"url_shortening/infra/db/postgres"
	"url_shortening/internal/useCase/urlShortening"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	app *fiber.App
	db  *postgres.Postgres
}

func NewServer(app *fiber.App, p *postgres.Postgres) (*Server, error) {

	return &Server{app: app, db: p}, nil
}

func (s *Server) Router() {
	s.app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	s.app.Post("/url/register", func(c fiber.Ctx) error {
		return urlShortening.Register(c, s.db)
	})
}
