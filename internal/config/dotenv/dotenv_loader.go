package dotenv

import (
	"fmt"

	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/joho/godotenv"
)

type dotenvLoader struct {
}

func (l dotenvLoader) Load() (*config.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load dotenv: %w", err)
	}

	cfg := &config.Config{}

	cfg.LogLevel, err = loadLogLevel()
	if err != nil {
		return nil, err
	}

	err = loadHttpConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = loadHandlerConfig(cfg)
	if err != nil {
		return nil, err
	}

	err = loadMinIOConfig(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func New() config.ConfigLoader {
	return dotenvLoader{}
}
