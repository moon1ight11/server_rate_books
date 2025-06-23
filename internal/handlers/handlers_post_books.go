package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rate_books/internal/database"
	"rate_books/internal/model"
)

// новая книга
func PostNewBook(c *gin.Context) {
	var Book model.Authors
	if err := c.ShouldBindJSON(&Book); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	database.InsertNewBook()

	c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
}
