package urlShortening

import (
	"encoding/json"
	"url_shortening/infra/db/postgres"
	"url_shortening/internal/domain/repository/urlShortening"

	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx, db *postgres.Postgres) error {

	type Request struct {
		Url string `json:"url"`
	}

	body := c.Body()

	var request Request

	if err := json.Unmarshal(body, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	repository := urlShortening.NewUrlShorteningRepository(db)

	err := repository.RegisterUrl(&request.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register url",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": request.Url,
	})
}
