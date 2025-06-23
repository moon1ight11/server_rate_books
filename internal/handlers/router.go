package handlers

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	// инициализация gin
	r := gin.Default()

	// get-роуты для книг
	r.GET("/all_books", GetAllBooks)
	r.GET("/ten_old_books", GetTenOldBooks)
	r.GET("/ten_new_books", GetTenNewBooks)
	r.GET("/max_rate_books", GetMAXBooks)
	r.GET("/min_rate_books", GetMINBooks)

	// get-роуты для авторов
	r.GET("/all_authors", GetAllAuthors)

	// post-роуты
	r.POST("/new_book", PostNewBook)
	r.POST("/new_author", PostNewAuthor)

	// запуск сервера
	r.Run(":8080")
}
