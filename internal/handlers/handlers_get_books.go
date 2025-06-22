package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"rate_books/internal/model"
)

// одна книга по номеру
func GetBookByID(c *gin.Context) {
	
}

// все книги
func GetAllBooks(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY rate DESC, author, year_public
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

// максимальные оценки
func GetMAXBooks(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							WHERE rate in (select MAX(rate) FROM rate_books)
							ORDER BY rate, author
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

// минимальные оценки
func GetMINBooks(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							WHERE rate in (select MIN(rate) FROM rate_books)
							ORDER BY rate, author
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}

// топ-10 старых
func GetTenOldBooks(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY year_public, title
							LIMIT 10
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

// топ-10 новых
func GetTenNewBooks(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY year_public DESC, title
							LIMIT 10
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

// топ по году прочтения
func GetTopBooksByYear(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	year_r := c.Param("year_read")

	rows, err := db.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							WHERE year_read = $1
							ORDER BY rate DESC, author
							`, year_r,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}
