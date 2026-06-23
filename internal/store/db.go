package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	goose "github.com/pressly/goose/v3"
)

type Store struct {
	db *sql.DB
}

func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		fmt.Println("Error while connecting DB", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error while connecting DB", err)
		return nil, err
	}

	return &Store{db: db}, nil
}

func (s Store) Migrate(dir string) error {
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("goose dialect: %w", err)
	}

	if err := goose.Up(s.db, dir); err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
