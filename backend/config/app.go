package config

import (
	"log"
	"os"
)

type AppConfig struct {
<<<<<<< HEAD
	// App
	Env  string
	Port string

	// Security
	JWTSecret string

	// CORS
	CORSOrigins string

	// Database
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	DBPort    string
	DBSSLMode string
	DBURL     string

	// Redis
	RedisAddr     string
	RedisPassword string

	// Bootstrap
	BootstrapAdminUser     string
	BootstrapAdminPassword string
=======
	Port        string
	CORSOrigins string
	JWTSecret   string
	DBHost      string
	DBUser      string
	DBPass      string
	DBName      string
	DBPort      string
	DBSSLMode   string
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
}

var App AppConfig

func LoadConfig() {
	App = AppConfig{
<<<<<<< HEAD
		Env:  getEnv("APP_ENV", "development"),
		Port: getEnv("APP_PORT", "8080"),

		JWTSecret: mustEnv("JWT_SECRET"),

		CORSOrigins: getEnv("CORS_ORIGINS", ""),

		DBHost:    mustEnv("DB_HOST"),
		DBUser:    mustEnv("DB_USER"),
		DBPass:    mustEnv("DB_PASSWORD"),
		DBName:    mustEnv("DB_NAME"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBSSLMode: getEnv("DB_SSLMODE", "require"),
		DBURL:     getEnv("DB_URL", ""),

		RedisAddr:     mustEnv("REDIS_ADDR"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		BootstrapAdminUser:     getEnv("BOOTSTRAP_ADMIN_USER", ""),
		BootstrapAdminPassword: getEnv("BOOTSTRAP_ADMIN_PASSWORD", ""),
=======
		Port:        getEnv("APP_PORT", "8080"),
		CORSOrigins: getEnv("CORS_ORIGINS", "http://localhost:5173"),
		JWTSecret:   mustEnv("JWT_SECRET"),
		DBHost:      mustEnv("DB_HOST"),
		DBUser:      mustEnv("DB_USER"),
		DBPass:      mustEnv("DB_PASSWORD"),
		DBName:      mustEnv("DB_NAME"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("%s is required", key)
	}
	return v
<<<<<<< HEAD
}
=======
}
>>>>>>> 3434ee2af12cefcd0083cdb7544b4f68644e7173
