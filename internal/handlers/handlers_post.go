package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"rate_books/internal/model"
)

// добавление новой книги
func AddNewBook(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	query := "INSERT INTO rate_books (author, title, rate, year_public, year_read) VALUES ($1, $2, $3, $4, $5)"

	_, err = db.Exec(query, book.Author.Author_name, book.Title, book.Rate, book.Year_public, book.Year_read)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Книга успешно добавлена"})
}

// добавление нового автора
func AddNewAuthor(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	var author model.Authors
	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON((http.StatusBadRequest), gin.H{"error": err.Error()})
	}

	query := "INSERT INTO authors (author,year_b,country) VALUES ($1, $2, $3)"

	_, err = db.Exec(query, author.Author_name, author.Year_born, author.Country)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Автор успешно добавлен"})
}
