package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rate_books/internal/database"
)

// список всех авторов
func GetAllAuthors(c *gin.Context) {
	all_authors, err := database.SelectAllAuthors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"all_authors": all_authors})
}
