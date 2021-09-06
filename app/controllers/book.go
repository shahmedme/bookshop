package controllers

import (
	"github.com/gofiber/fiber/v2"
	"shakilahmed.com/bookshop/app/models"
	"shakilahmed.com/bookshop/platform/database"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DB
	var books []models.Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book models.Book
	db.Find(&book, id)
	return c.JSON(book)
}

func AddBook(c *fiber.Ctx) error {
	db := database.DB
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var bookModel models.Book
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	db.First(&bookModel, id)

	if bookModel.Title == "" {
		return c.Status(500).SendString("No book found with this ID")
	}

	bookModel.Title = book.Title
	bookModel.Author = book.Author
	bookModel.Rating = book.Rating

	db.Save(&bookModel)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var book models.Book

	db.First(&book, id)
	db.Delete(&book)

	return c.SendString("Book successfully deleted")
}

func Test(c *fiber.Ctx) error {
	return c.SendString("Test success...........")
}
