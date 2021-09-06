package routes

import "github.com/gofiber/fiber/v2"

func AdminRoutes(a *fiber.App) {
	a.Get("/admin", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Name": "Shakil Ahmed",
		})
	})
}
