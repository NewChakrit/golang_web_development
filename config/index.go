package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Config envConfig

type envConfig struct {
	AppPort string
	//DBPort  string
}

func init() {
	Config.LoadConfig()
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	e.AppPort = loadString("APP_PORT", ":8080")
	//e.DBPort = loadString("DB_PORT", ":8080")
}

func loadString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panic("APP_PORT not set in .env file")
		return fallback
	}

	return val
}
