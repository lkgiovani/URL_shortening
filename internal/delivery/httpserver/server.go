package httpserver

import (
	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/useCase/urlShortening"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App    *fiber.App
	Db     *postgres.Postgres
	Redis  *redis.Redis
	Config *environment.Config
}

func NewServer(app *fiber.App, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) (*Server, error) {

	return &Server{App: app, Db: db, Redis: redis, Config: config}, nil
}

func (s *Server) Router() {
	s.App.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	s.App.Post("/url/register", func(c fiber.Ctx) error {
		return urlShortening.Register(c, s.Db, s.Redis, s.Config) // TODO: change to useCase
	})

	s.App.Get("/:urlShortened", func(c fiber.Ctx) error {
		return urlShortening.GetUrl(c, s.Db, s.Redis, s.Config)
	})
}
