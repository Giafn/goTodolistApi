package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env        string
	Host       string
	Port       string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		Env:        os.Getenv("ENV"),
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_NAME"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
	}

	return config
}
