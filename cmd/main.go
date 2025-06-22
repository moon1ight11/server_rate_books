package main

import (
	"rate_books/internal/database"
	"rate_books/internal/handlers"
)

func main() {
	// подключение к базе данных, обновление миграций
	database.UppingMigrations()

	// роутер
	handlers.Router()
}