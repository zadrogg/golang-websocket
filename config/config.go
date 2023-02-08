package config

import (
	"os"
	"strconv"
)

type Config struct {
	Sentry dashSentry
	Redis  redisDB
	Cache  cacheValue
	Server server
}

type dashSentry struct {
	Dsn string
}

type redisDB struct {
	Host     string
	Port     int
	Password string
}

type cacheValue struct {
	Sign        bool
	TypeStorage string
}

type server struct {
	Url string
}

func GetConfig() *Config {
	return &Config{
		Redis: redisDB{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
		},
		Cache: cacheValue{
			Sign:        getEnvBool("CACHE", false),
			TypeStorage: getEnv("CACHE_TYPE", "memory"),
		},
		Server: server{
			Url: getEnv("SERVER", "localhost:3000"),
		},
		Sentry: dashSentry{
			Dsn: getEnv("SENTRY_DSN", ""),
		},
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultValue
}
