package utils

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvironmentVarExists(key string) string {
	payload, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("PLATFORM environment variable not set for %s", key)
	}
	return payload
}

func OpenEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func OpenDB(dbType string, dbURL string) *sql.DB {
	db, err := sql.Open(dbType, dbURL)
	if err != nil {
		log.Printf("Unable to Open DB, check ports and connecion string")
		return nil
	}
	return db
}
