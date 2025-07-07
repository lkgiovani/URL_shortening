package urlShortening

import (
	"encoding/json"
	"fmt"
	"time"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/urlShortening_repo"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Url string `json:"url" validate:"required,url"`
}

func Register(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {

	body := c.Body()

	var request RegisterRequest

	err := json.Unmarshal(body, &request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(request)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Validation error: %s", errors),
		})
	}

	if err = json.Unmarshal(body, &request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	repository := urlShortening_repo.NewUrlShorteningRepository(db, config)

	urlShortened, err := repository.RegisterUrl(&request.Url, c.Locals("id").(string))
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

	c.Redirect(urlOriginal.UrlOriginal, 302)
	return nil
}
