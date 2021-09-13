package controllers

import "github.com/gofiber/fiber/v2"

func AdminHome(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"name": "Shakil Ahmed",
	}, "layouts/main")
}
