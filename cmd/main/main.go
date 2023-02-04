package main

import (
	"learnpackage/pkg/controllers/createBooks"
	"learnpackage/pkg/controllers/getBooks"
	"learnpackage/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}


func main() {
	validate = validator.New()
	router := gin.Default()

	router.GET("/books" , getBooks.GetAllBooks)
	router.POST("/books" , createBooks.CreateBook)
	router.GET("/books/:id" , getBooks.GetBooksById)


	router.Run("localhost:8080")
}