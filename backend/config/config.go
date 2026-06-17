package config

import (
	"os"
)

type Config struct {
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	RedisAddr   string
	JWTSecret   string
	Port        string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "opencti"),
		DBUser:     getEnv("DB_USER", "opencti"),
		DBPassword: getEnv("DB_PASSWORD", "opencti"),
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:  getEnv("JWT_SECRET", "opencti-secret-change-in-prod"),
		Port:       getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
