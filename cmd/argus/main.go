package main

import (
	"fmt"

	"github.com/elton451/argus/internal/store"
)

func main() {
	config := store.Load()

	fmt.Println(config);

	store, err := store.Open(config.DBPath)
	if err != nil {
		fmt.Println("Error while connecting DB", err)
	}

	migrateErr := store.Migrate("internal/migrations")
	if migrateErr != nil {
		fmt.Println("Error while on migration", migrateErr)
	}
}

