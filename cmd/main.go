package main

import (
	// "fmt"
	// "os"
	"rate_books/internal/database"
	"rate_books/internal/handlers"
)

func main() {
	// подключение к базе данных, обновление миграций
	database.UppingMigrations()

	handlers.Router()


}