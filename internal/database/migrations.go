package database

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
)

func UppingMigrations() {
	const (
		url           = "postgres://fedor:fedor_pass@localhost:15432/rate"
		migrationsDir = "./migrations"
	)

	config, err := pgx.ParseConfig(url)
	if err != nil {
		log.Fatalf("Failed to parse database URL: %v", err)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	goose.SetBaseFS(nil)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	err = goose.Up(db, migrationsDir)
	if err != nil {
		log.Fatalf("Failed to up migrations: %v", err)
	}
}
