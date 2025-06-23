package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rate_books/internal/database"
)

// список всех книг
func GetAllBooks(c *gin.Context) {
	all_books, err := database.SelectAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"all_books": all_books})
}

// список книг с максимальной оценкой
func GetMAXBooks(c *gin.Context) {
	max_books, err := database.SelectMAXBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"max_books": max_books})
}

// список книг с максимальной оценкой
func GetMINBooks(c *gin.Context) {
	min_books, err := database.SelectMINBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"min_books": min_books})
}

// топ-10 старых
func GetTenOldBooks(c *gin.Context) {
	old_books, err := database.SelectTopOldBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"old_books": old_books})
}

// топ-10 новых
func GetTenNewBooks(c *gin.Context) {
	new_books, err := database.SelectTopNewBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"new_books": new_books})
}
