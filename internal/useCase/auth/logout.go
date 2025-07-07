package auth

import (
	"time"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {
	// Limpar o cookie do token
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour) // Expirar no passado
	cookie.HTTPOnly = true
	cookie.Secure = false
	cookie.SameSite = "Lax"
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout successful",
	})
}
