package urlShortening

import (
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/urlShortening_repo"

	"github.com/gofiber/fiber/v2"
)

func ListUserUrls(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {
	// Obter ID do usuário do contexto (via middleware de autenticação)
	userID, ok := c.Locals("id").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not authenticated",
		})
	}

	repository := urlShortening_repo.NewUrlShorteningRepository(db, config)
	urls, err := repository.GetUserUrls(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve URLs",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"urls": urls,
	})
}
