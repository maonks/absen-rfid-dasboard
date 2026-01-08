package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeviceRoute(app *fiber.App, db *gorm.DB) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	app.Post("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
}
