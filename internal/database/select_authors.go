package database

import (
	"rate_books/internal/model"
)

// запрос на всех авторов из списка
func SelectAllAuthors() ([]model.Authors, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT author, year_b, country
							FROM authors
							ORDER BY author
							`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Authors []model.Authors
	for rows.Next() {
		var Author model.Authors
		err := rows.Scan(&Author.Author_name, &Author.Year_born, &Author.Country)
		if err != nil {
			return nil, err
		}
		Authors = append(Authors, Author)
	}

	return Authors, nil
}
