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
	return &Server{App: app, Db: db, Redis: redis, Config: config}, nil
}

// URL handlers
func (s *Server) handleURLRegister(c *fiber.Ctx) error {
	return urlShortening.Register(c, s.Db, s.Redis, s.Config)
}

func (s *Server) handleURLGet(c *fiber.Ctx) error {
	return urlShortening.GetUrl(c, s.Db, s.Redis, s.Config)
}

// Auth handlers
func (s *Server) handleAuthRegister(c *fiber.Ctx) error {
	return auth.Register(c, s.Db, s.Redis, s.Config)
}

func (s *Server) handleAuthLogin(c *fiber.Ctx) error {
	return auth.Login(c, s.Db, s.Redis, s.Config)
}

// Home handler
func (s *Server) handleHome(c *fiber.Ctx) error {
	return c.SendString("salve! 🤙")
}

func (s *Server) Router() {
	// Home
	s.App.Get("/", s.handleHome)

	authGroup := s.App.Group("/auth")

	authGroup.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 1 * time.Minute,
	}))

	authGroup.Post("/register", s.handleAuthRegister)
	authGroup.Post("/login", s.handleAuthLogin)

	// Só para /url/register
	s.App.Use("/register", limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	}))

	s.App.Use("/register", func(c *fiber.Ctx) error {
		return middleware.AuthMiddleware(c, s.Config)
	})

	s.App.Post("/register", s.handleURLRegister)
	s.App.Get("/:urlShortened", s.handleURLGet)

}
