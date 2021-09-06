package routes

import (
	"github.com/gofiber/fiber/v2"
	"shakilahmed.com/bookshop/app/controllers"
)

func AccountRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1")

	v1.Get("/hello", controllers.HelloWorld)
}
