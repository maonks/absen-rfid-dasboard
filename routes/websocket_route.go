package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/maonks/absen-rfid-backend/controllers"
)

func WebSocketRoute(app *fiber.App) {
	app.Get("/websocket", websocket.New(controllers.WebsocketHandler))
}
