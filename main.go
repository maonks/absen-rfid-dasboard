package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/maonks/absen-rfid-backend/configs"
	"github.com/maonks/absen-rfid-backend/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println(".ENV Tidak di temukan")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db, err := configs.KonekDB()

	if err != nil {
		log.Fatal("Gagal Konek DB", err)
	}

	routes.DeviceRoute(app, db)
	routes.WebSocketRoute(app)
	routes.AbsenRoute(app, db)

	app.Listen("" + os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT") + "")
}
