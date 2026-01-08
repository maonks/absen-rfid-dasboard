package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maonks/absen-rfid-backend/controllers"
	"gorm.io/gorm"
)

func AbsenRoute(app *fiber.App, db *gorm.DB) {

	app.Post("/api/absen", controllers.CreateAbsen(db))
	app.Get("/api/absen", controllers.SearchAbsen(db))
	app.Put("/api/kartu/:uid", controllers.UpdateKartu(db))

}
