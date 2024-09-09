package config

import (
	"os"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPasswd   string
	DBName     string
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "3000"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPasswd:   getEnv("DB_PASSWD", "password"),
		DBName:     getEnv("DB_NAME", "ecomgo"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
