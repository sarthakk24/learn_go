package createBooks

import (
	"fmt"
	"learnpackage/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


var validate *validator.Validate

var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func CreateBook(c *gin.Context) {
	
	validate = validator.New()
	var newBook models.Book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	validationErr := validate.Struct(newBook)


	if validationErr != nil {
		fmt.Print(validationErr)
		return
	}
	books = append(books , newBook)
	c.IndentedJSON(http.StatusOK , newBook)
}