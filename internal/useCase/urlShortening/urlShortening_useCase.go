package urlShortening

import (
	"encoding/json"
	"time"

	"url_shortening/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/urlShortening_repo"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {

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

	repository := urlShortening_repo.NewUrlShorteningRepository(db, config)

	urlShortened, err := repository.RegisterUrl(&request.Url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = redis.Set(urlShortened.Slug, request.Url, 3*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set url in redis",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": urlShortened.UrlShortened,
	})
}

func GetUrl(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {
	urlShortened := c.Params("urlShortened")

	url, err := redis.Get(urlShortened)
	if err == nil {
		c.Redirect(url, 302)
		return nil
	}

	repository := urlShortening_repo.NewUrlShorteningRepository(db, config)
	urlOriginal, err := repository.GetUrl(urlShortened)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "URL not found",
		})
	}

	err = redis.Set(urlOriginal.Slug, urlOriginal.UrlOriginal, 3*time.Minute)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set url in redis",
		})
	}

	c.Redirect(url, 302)
	return nil
}
