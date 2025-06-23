package database

import (
	"rate_books/internal/model"
)

// запрос на добавление нового автора
func InsertNewAuthor(a model.Authors) error {
	query := "INSERT INTO authors (author,year_b,country) VALUES ($1, $2, $3)"

	_, err := DB.Exec(query, a.Author_name, a.Year_born, a.Country)
	if err != nil {
		return err
	}
	return nil
}
