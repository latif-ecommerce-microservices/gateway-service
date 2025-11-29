package config

import "os"

type Config struct {
	Port         string
	UserService  string
	OrderService string
}

func LoadConfig() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		UserService:  getEnv("USER_SERVICE", "http://localhost:9001"),
		OrderService: getEnv("ORDER_SERVICE", "http://localhost:9002"),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
