package httpserver

import (
	"time"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/delivery/httpserver/middleware"
	"url_shortening/internal/useCase/auth"
	"url_shortening/internal/useCase/urlShortening"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

type Server struct {
	App    *fiber.App
	Db     *postgres.Postgres
	Redis  *redis.Redis
	Config *environment.Config
}

func NewServer(app *fiber.App, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) (*Server, error) {

	app.Use("/auth", limiter.New(
		limiter.Config{
			Max:        20,
			Expiration: 1 * time.Minute,
		},
	))

	app.Use("/url/register", limiter.New(
		limiter.Config{
			Max:        100,
			Expiration: 1 * time.Minute,
		},
	))

	return &Server{App: app, Db: db, Redis: redis, Config: config}, nil
}

func (s *Server) Router() {
	s.App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("salve! 🤙")
	})

	s.App.Post("/url/register", func(c *fiber.Ctx) error {
		return middleware.AuthMiddleware(c, s.Config)
	}, func(c *fiber.Ctx) error {
		return urlShortening.Register(c, s.Db, s.Redis, s.Config) // TODO: change to useCase
	})

	s.App.Get("/:urlShortened", func(c *fiber.Ctx) error {
		return urlShortening.GetUrl(c, s.Db, s.Redis, s.Config)
	})

	s.App.Post("/auth/register", func(c *fiber.Ctx) error {
		return auth.Register(c, s.Db, s.Redis, s.Config)
	})

	s.App.Post("/auth/login", func(c *fiber.Ctx) error {
		return auth.Login(c, s.Db, s.Redis, s.Config)
	})
}
