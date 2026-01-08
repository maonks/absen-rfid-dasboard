package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maonks/absen-rfid-backend/models"
	"gorm.io/gorm"
)

func UpdateKartu(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var k struct{ Nama string }
		uid := c.Params("uid")

		c.BodyParser(&k)

		db.Model(&models.Kartu{}).Where("uid=?", uid).Update("nama", k.Nama)
		return c.SendStatus(200)
	}
}
