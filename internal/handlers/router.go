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

	// маршруты
	// r.GET("/all_rate", GetAllBooks)  							// запрос на все книги
	// r.GET("/ten_old_books", GetTenOldBooks) 						// запрос на 10 старых книг
	// r.GET("/top_books_by_year/:year_read", GetTopBooksByYear) 	// запрос на 10 книг по году прочтения
	// r.GET("/ten_new_books", GetTenNewBooks) 						// запрос на 10 новых книг
	// r.GET("/max_books", GetMAXBooks)								// запрос на максимальный рейтинг
	// r.GET("/min_books", GetMINBooks)								// запрос на минимальный рейтинг

	// r.POST("/add_new_book", AddNewBook)							// добавление новой книги

	r.Run(":8080")
}
