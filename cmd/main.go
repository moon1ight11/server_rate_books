package main

import (
	"rate_books/internal/database"
	"rate_books/internal/handlers"
)

func main() {
	// подключение к базе данных
	database.Connection()
	defer database.DB.Close()

	// обновление миграций
	database.UppingMigrations()

	// роутер
	handlers.Router()
}
