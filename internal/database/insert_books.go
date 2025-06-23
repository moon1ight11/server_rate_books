package database

import (
	"rate_books/internal/model"
)

// запрос на добавление новой книги
func InsertNewBook() error {
	defer DB.Close()
	var Book model.Book
	query := "INSERT INTO rate_books (author, title, rate, year_public, year_read) VALUES ($1, $2, $3, $4, $5)"

	_, err := DB.Exec(query, Book.Author.Author_name, Book.Title, Book.Rate, Book.Year_public, Book.Year_read)
	if err != nil {
		return err
	}
	return nil
}