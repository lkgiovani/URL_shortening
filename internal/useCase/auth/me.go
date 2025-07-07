package auth

import (
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/user_repo"
	"url_shortening/pkg/jwtpkg"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {
	// Obter token do cookie
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token not found",
		})
	}

	// Verificar e decodificar o token
	claims, err := jwtpkg.ValidateToken(token, config.JWT_SECRET)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Extrair email dos claims
	emailClaim, ok := claims["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Buscar usuário no banco
	repository := user_repo.NewUserRepository(db, config)
	user, err := repository.GetUserByEmail(emailClaim)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
