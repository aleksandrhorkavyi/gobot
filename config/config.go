package config

import (
	"os"
)

type BotConfig struct {
	Url string
}

type Config struct {
	Bot BotConfig
}

func (c *Config) BotUrl() string {
	return c.Bot.Url
}

func New() *Config {
	return &Config{
		Bot: BotConfig{
			Url: getEnv("BOT_URL", ""),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}