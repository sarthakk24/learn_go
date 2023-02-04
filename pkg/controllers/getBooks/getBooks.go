package getBooks

import (
	"errors"
	"learnpackage/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


var books = []models.Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

// Get All Books
func GetAllBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK , books)
} 

// Get Books By ID
func booksById(id string) (*models.Book, error) {
	for i,element := range books{
		if element.ID == id {
			return &books[i] , nil
		}
	}
	return nil, errors.New("No book was found")
}

func GetBooksById(c *gin.Context) {
	id := c.Param("id")
	foundBook , error  := booksById(id)
	if error != nil {
		c.IndentedJSON(http.StatusNotFound  , gin.H{"Message" : error.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK , foundBook)
}

