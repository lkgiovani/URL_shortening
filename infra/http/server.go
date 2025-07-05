package http

import "github.com/gofiber/fiber/v3"

func NewServer(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
