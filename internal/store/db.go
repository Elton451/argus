package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
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

