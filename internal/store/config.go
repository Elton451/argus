package store

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath string
	Addr   string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPath := os.Getenv("ARGUS_DB_PATH")
	addr := os.Getenv("ARGUS_ADDR")

	return Config{
		DBPath: getEnv(dbPath, "./argus.db"),
		Addr:   getEnv(addr, ":8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
