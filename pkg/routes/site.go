package routes

import (
	"github.com/gofiber/fiber/v2"
	"shakilahmed.com/bookshop/app/controllers"
)

func SiteRoute(a *fiber.App) {
	a.Get("/", controllers.Home)
}
