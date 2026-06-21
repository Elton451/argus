package store

import "os"

type Config struct {
    DBPath string
    Addr   string
}

func Load() Config {
    return Config{
        DBPath: getEnv("ARGUS_DB_PATH", "./argus.db"),
        Addr:   getEnv("ARGUS_ADDR", ":8080"),
    }
}

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}
