package config

import (
	"os"
)

type RedisConfig struct {
	Addr      string
	KeyPrefix string
	Network   string
}

type UsersConfig struct {}

type Config struct {
	ServerAddr string
	Redis      *RedisConfig
}

func New() *Config {
	return &Config{
		ServerAddr: optionalEnv("SERVER_ADDR", ":3001"),

		Redis: &RedisConfig{
			Addr:      optionalEnv("REDIS_ADDR", "127.0.0.1:6379"),
			KeyPrefix: optionalEnv("REDIS_KEY_PREFIX", "go-api::"),
			Network:   optionalEnv("REDIS_NETWORK", "tcp"),
		},
	}
}

func optionalEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
