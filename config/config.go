package config

import (
	"os"
)

func New() *Config {
	return &Config{
		Bot: BotConfig{
			Url: getEnv("BOT_URL", ""),
		},
		Db: DbConfig{
			MigrationsDirectory: "db/migrations/",
		},
	}
}

type BotConfig struct {
	Url string
}

type DbConfig struct {
	MigrationsDirectory string
}

type Config struct {
	Bot BotConfig
	Db DbConfig
}

func (c *Config) BotUrl() string {
	return c.Bot.Url
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
