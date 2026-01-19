package config

import (
	"strings"

	"github.com/latif-ecommerce-microservices/gateway-service/pkg/envparser"
	"github.com/latif-ecommerce-microservices/gateway-service/pkg/logging"
)

type Config struct {
	AppHost      string `env:"APP_HOST"`
	AppHTTPPort  string `env:"APP_HTTP_PORT"`
	UserGRPCPort string `env:"USER_GRPC_PORT"`
	JWTSecret    string `env:"JWT_SECRET"`
	LogLevel     string `env:"LOG_LEVEL"`
	Redis        RedisConfig
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
	DB       int    `env:"REDIS_DB"`
}

func (c *Config) GetLogLevel() logging.Level {
	switch strings.ToLower(c.LogLevel) {
	case "debug":
		return logging.LevelDebug
	case "info":
		return logging.LevelInfo
	case "warn":
		return logging.LevelWarn
	case "error":
		return logging.LevelError
	case "fatal":
		return logging.LevelFatal
	default:
		return logging.LevelInfo
	}
}

func GetConfig() (*Config, error) {
	cfg := Config{}
	err := envparser.LoadEnv(&cfg)
	if err != nil {
		return nil, nil
	}

	return &cfg, nil
}
