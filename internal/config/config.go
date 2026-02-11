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

	MinIO struct {
		Endpoint        string
		AccessKeyID     string
		SecretAccessKey string
		UseSSL          bool
	}

	CouchDB struct {
		Port           int
		Hostname       string
		Username       string
		Password       string
		Insecure       bool
		MetadataDBName string
	}
}

type ConfigLoader interface {
	Load() (*Config, error)
}
