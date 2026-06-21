package main

import (
	"fmt"

	"github.com/elton451/argus/internal/store"
)

func main() {
	config := store.Load()

	fmt.Println(config);

	db, err := store.Open(config.DBPath)
	if err != nil {
		fmt.Println("Error while connecting DB", err)
	}

	fmt.Println("DB: ", db)
}

