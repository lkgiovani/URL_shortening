package urlShortening

import (
	"encoding/json"

	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/urlShortening"

	"github.com/gofiber/fiber/v3"
)

func Register(c fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {

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

	repository := urlShortening.NewUrlShorteningRepository(db, config)

	urlShortened, slug, err := repository.RegisterUrl(&request.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to register url",
		})
	}

	err = redis.Set(slug[:8], request.Url, 0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set url in redis",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": urlShortened,
	})
}

func GetUrl(c fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {

	urlShortened := c.Params("urlShortened")

	url, err := redis.Get(urlShortened)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	c.Redirect().Status(302).To(url)
	return nil
}
