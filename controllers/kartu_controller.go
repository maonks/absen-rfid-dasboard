package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maonks/absen-rfid-backend/models"
	"gorm.io/gorm"
)

func UpdateKartu(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req struct {
			Nama string `form:"nama"`
		}

		uid := c.Params("uid")

		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString("Invalid form data")
		}

		if req.Nama == "" {
			return c.Status(422).SendString("Nama tidak boleh kosong")
		}

		result := db.Model(&models.Kartu{}).
			Where("uid = ?", uid).
			Update("nama", req.Nama)

		if result.Error != nil {
			return c.Status(500).SendString("Gagal update nama")
		}

		if result.RowsAffected == 0 {
			return c.Status(404).SendString("Kartu tidak ditemukan")
		}

		return c.SendStatus(200)
	}
}
