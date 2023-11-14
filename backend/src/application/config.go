package application

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort   uint16
	RedisAddress string
}

func LoadConfig() Config {
	cfg := Config{
		RedisAddress: "127.0.0.1:6379",
		ServerPort:   8080,
	}

	if redisAddress, exists := os.LookupEnv("REDIS_ADDRESS"); exists {
		cfg.RedisAddress = redisAddress
	}

	if port, exists := os.LookupEnv("PORT"); exists {
		if portInt, err := strconv.Atoi(port); err == nil {
			cfg.ServerPort = uint16(portInt)
		}
	}

	return cfg
}
