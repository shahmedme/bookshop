package routes

import (
	"github.com/gofiber/fiber/v2"
	"shakilahmed.com/bookshop/app/controllers"
)

func BookRoutes(a *fiber.App) {
	v1 := a.Group("/api/v1/book")

	v1.Get("", controllers.GetBooks)
	v1.Get("/:id", controllers.GetBook)
	v1.Post("", controllers.AddBook)
	v1.Put("/:id", controllers.UpdateBook)
	v1.Delete("/:id", controllers.DeleteBook)
	v1.Get("/test", controllers.Test)
}
