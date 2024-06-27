package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort int
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("NOTES_PORT"))
	if err != nil {
		log.Fatalf("Error converting SERVER_PORT to integer: %v", err)
	}

	AppConfig = Config{
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		JWTSecret:  os.Getenv("SECRET_KEY"),
		ServerPort: port,
	}
}
