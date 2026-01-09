package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func EditModal(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid := c.Params("uid")

		return c.Render("edit_modal", fiber.Map{
			"Uid": uid,
		})
	}
}
