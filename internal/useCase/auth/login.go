package auth

import (
	"encoding/json"
	"fmt"
	"time"
	"url_shortening/infra/config/environment"
	"url_shortening/infra/db/postgres"
	"url_shortening/infra/db/redis"
	"url_shortening/internal/domain/repository/user_repo"
	"url_shortening/pkg/cryptPkg"
	"url_shortening/pkg/jwtpkg"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func Login(c *fiber.Ctx, db *postgres.Postgres, redis *redis.Redis, config *environment.Config) error {

	body := c.Body()

	var user LoginRequest

	err := json.Unmarshal(body, &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}

	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(user)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Validation error: %s", errors),
		})
	}

	repository := user_repo.NewUserRepository(db, config)
	response, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !cryptPkg.ComparePassword(user.Password, response.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "email or password is incorrect",
		})
	}

	token, err := jwtpkg.GenerateToken(jwt.MapClaims{
		"id":    response.ID,
		"email": user.Email,
	}, config.JWT_SECRET)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = true
	cookie.Secure = false // Set to true in production with HTTPS
	cookie.SameSite = "Lax"
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"user": fiber.Map{
			"id":    response.ID,
			"name":  response.Name,
			"email": response.Email,
		},
	})
}
