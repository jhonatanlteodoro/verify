// Package responsible for hold the app environment vars, envery env var should be
// defined here
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_DRIVER       string
	MYSQL_USERNAME  string
	MYSQL_PASSWORD  string
	MYSQL_DB_NAME   string
	MYSQL_HOST      string
	MYSQL_HOST_PORT string
	SQLITE_FILENAME string
	ENV             string
}

func loadEnvVars() {
	log.Println("Loading ENV vars..")
	dotenvFilePath, envfileErr := filepath.Abs("./.env")
	if envfileErr != nil {
		log.Fatal(envfileErr)
	}

	log.Printf("Loading ENV file from: %s\n", dotenvFilePath)

	err := godotenv.Load(dotenvFilePath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("ENV vars loaded...")
}

func InitConfig() *Config {
	loadEnvVars()
	return &Config{
		DB_DRIVER:       os.Getenv("DB_IN_USE"),
		MYSQL_DB_NAME:   os.Getenv("MYSQL_DB_NAME"),
		MYSQL_HOST:      os.Getenv("MYSQL_HOST"),
		MYSQL_USERNAME:  os.Getenv("MYSQL_USERNAME"),
		MYSQL_PASSWORD:  os.Getenv("MYSQL_PASSWORD"),
		MYSQL_HOST_PORT: os.Getenv("MYSQL_HOST_PORT"),
		SQLITE_FILENAME: os.Getenv("SQLITE_FILENAME"),
	}
}
