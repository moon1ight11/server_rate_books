package database

import (
	"rate_books/internal/model"
)

// запрос на все книги из списка
func SelectAllBooks() ([]model.Book, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY rate DESC, author, year_public
							`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на список книг с максимальной оценкой
func SelectMAXBooks() ([]model.Book, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							WHERE rate in (select MAX(rate) FROM rate_books)
							ORDER BY rate, author
							`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на список книг с минимальной оценкой
func SelectMINBooks() ([]model.Book, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							WHERE rate in (select MIN(rate) FROM rate_books)
							ORDER BY rate, author
							`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на топ-10 старых книг
func SelectTopOldBooks() ([]model.Book, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY year_public, title
							LIMIT 10
							`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}

// запрос на топ-10 новых книг
func SelectTopNewBooks() ([]model.Book, error) {
	defer DB.Close()
	rows, err := DB.Query(`SELECT title, author, year_public, year_read, rate
							FROM rate_books
							ORDER BY year_public DESC, title
							LIMIT 10
							`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var Books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.Title, &book.Author.Author_name, &book.Year_public, &book.Year_read, &book.Rate)
		if err != nil {
			return nil, err
		}
		Books = append(Books, book)
	}
	return Books, nil
}
