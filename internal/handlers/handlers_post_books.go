package handlers

import (
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

	err := database.InsertNewBook(Book) 
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
}
