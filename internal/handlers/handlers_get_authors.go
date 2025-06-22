package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"rate_books/internal/model"
)

// все авторы
func GetAllAuthors(c *gin.Context) {
	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	rows, err := db.Query(`SELECT author, year_b, country
							FROM authors
							ORDER BY author
							`,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var Authors []model.Authors
	for rows.Next() {
		var Author model.Authors
		err := rows.Scan(&Author.Author_name, &Author.Year_born, &Author.Country)
		if err != nil {
			log.Fatal(err)
		}
		Authors = append(Authors, Author)
	}

	c.JSON(http.StatusOK, gin.H{"authors": Authors})

}
