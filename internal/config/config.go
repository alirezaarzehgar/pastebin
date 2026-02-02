package config

import "github.com/alirezaarzehgar/pastebin/internal/logger"

type Config struct {
	LogLevel logger.Leveler
}

type ConfigLoader interface {
	Load() (*Config, error)
}
