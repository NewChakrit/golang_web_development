package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Config envConfig

type envConfig struct {
	AppPort            string
	DBPath             string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	JwtSaltKey         string
	FEOriginUrl        string
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
	e.DBPath = loadString("DB_PATH", "postgres://postgres:P@ssw0rd@localhost:5433/tasks?sslmode=disable")
	e.GoogleClientID = loadString("GOOGLE_CLIENT_ID", "")
	e.GoogleClientSecret = loadString("GOOGLE_CLIENT_SECRET", "")
	e.GoogleRedirectURL = loadString("GOOGLE_REDIRECT_URL", "")
	e.JwtSaltKey = loadString("JWT_SALT_KEY", "hi_test_salt")
	e.FEOriginUrl = loadString("FE_ORIGIN_URL", "https://react-login-beta-seven.vercel.app")
}

func loadString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panic("APP_PORT not set in .env file")
		return fallback
	}

	return val
}
