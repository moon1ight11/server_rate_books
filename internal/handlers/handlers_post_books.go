package handlers

import (
	"fmt"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"

	"github.com/gin-gonic/gin"
)

// новая книга
func PostNewBook(c *gin.Context) {
	var Book model.Book
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}
	fmt.Println(Book)

	database.InsertNewBook(Book) 

	c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
}
