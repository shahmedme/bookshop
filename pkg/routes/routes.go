package routes

import "github.com/gofiber/fiber/v2"

func InitRoutes(app *fiber.App) {
	AccountRoutes(app)
	AdminRoutes(app)
	BookRoutes(app)
}
