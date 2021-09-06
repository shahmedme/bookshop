package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django"
	"github.com/joho/godotenv"
	"shakilahmed.com/bookshop/pkg/routes"
	"shakilahmed.com/bookshop/platform/database"
)

func main() {
	err := godotenv.Load("./pkg/configs/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	templateEngine := django.New("./app/templates", ".html")
	app := fiber.New(fiber.Config{
		Views: templateEngine,
	})

	routes.InitRoutes(app)
	database.InitDatabase()
	log.Fatal(app.Listen(":8000"))
}
