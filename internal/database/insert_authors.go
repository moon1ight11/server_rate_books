package database

import (
	"rate_books/internal/model"
)

// запрос на добавление нового автора
func InsertNewAuthor() error {
	defer DB.Close()
	var Author model.Authors
	query := "INSERT INTO authors (author,year_b,country) VALUES ($1, $2, $3)"

	_, err := DB.Exec(query, Author.Author_name, Author.Year_born, Author.Country)
	if err != nil {
		return err
	}
	return nil
}
