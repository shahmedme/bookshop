package routes

import (
	"github.com/gofiber/fiber/v2"
	"shakilahmed.com/bookshop/app/controllers"
)

func AdminRoutes(a *fiber.App) {
	a.Get("/admin", controllers.AdminHome)
}
