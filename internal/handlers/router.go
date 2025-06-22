package handlers

import (
	"github.com/gin-gonic/gin"
)

const (
	url           = "postgres://fedor:fedor_pass@localhost:15432/rate"
	migrationsDir = "./migrations"
)

func Router() {
	// инициализация gin
	r := gin.Default()

	r.GET("/book_by_id/:id", GetBookByID)
	r.GET("/all_books", GetAllBooks)
	r.GET("/top_books_by_year/:year_read", GetTopBooksByYear)
	r.GET("/ten_old_books", GetTenOldBooks)
	r.GET("/ten_new_books", GetTenNewBooks)
	r.GET("/max_rate_books", GetMAXBooks)
	r.GET("/min_rate_books", GetMINBooks)

	r.GET("/all_authors", GetAllAuthors)

	r.POST("/new_book", AddNewBook)
	r.POST("/new_author", AddNewAuthor)
	
	r.Run(":8080")
}
