package middleware

import (
	"url_shortening/config/environment"
	"url_shortening/pkg/jwtpkg"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx, config *environment.Config) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	claims, err := jwtpkg.ValidateToken(token, config.JWT_SECRET)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	c.Locals("email", claims["email"])

	return c.Next()
}
