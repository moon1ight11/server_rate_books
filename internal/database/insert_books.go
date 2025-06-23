package database

import (
	"rate_books/internal/model"
)

// запрос на добавление новой книги
func InsertNewBook(b model.Book) error {
	query := "INSERT INTO rate_books (author, title, rate, year_public, year_read) VALUES ($1, $2, $3, $4, $5)"

	_, err := DB.Exec(query, b.Author, b.Title, b.Rate, b.Year_public, b.Year_read)
	if err != nil {
		return err
	}
	return nil
}