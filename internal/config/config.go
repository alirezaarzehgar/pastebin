package config

import (
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
)

type Config struct {
	LogLevel logger.Leveler

	Http struct {
		Address string
		Port    int

		IdleTimeout  time.Duration
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Handler struct {
		MaxFileSize int64
	}
}

type ConfigLoader interface {
	Load() (*Config, error)
}
