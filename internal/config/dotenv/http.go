package dotenv

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/alirezaarzehgar/pastebin/internal/config"
)

const (
	DefaultHttpAddress      = "0.0.0.0"
	DefaultHttpPort         = 8080
	DefaultHttpIdleTimeout  = 120
	DefaultHttpReadTimeout  = 120
	DefaultHttpWriteTimeout = 120
)

func parseInt(key string, defaultValue int) (int, error) {
	var value int
	valueString := os.Getenv(key)
	if valueString == "" {
		value = defaultValue
		return value, nil
	}
	value, err := strconv.Atoi(valueString)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s: %w", valueString, err)
	}

	return value, nil
}

func parseDuration(key string, defaultValue int) (time.Duration, error) {
	value, err := parseInt(key, defaultValue)
	if err != nil {
		return 0, err
	}

	duration := time.Duration(value) * time.Second
	return duration, nil
}

func loadHttpConfig(cfg *config.Config) (*config.Config, error) {
	addr := os.Getenv("HTTP_ADDRESS")
	if addr == "" {
		addr = DefaultHttpAddress
	}
	cfg.Http.Address = addr

	var err error
	cfg.Http.Port, err = parseInt("HTTP_PORT", DefaultHttpPort)
	if err != nil {
		return nil, err
	}

	cfg.Http.IdleTimeout, err = parseDuration("HTTP_IDLE_TIMEOUT", DefaultHttpIdleTimeout)
	if err != nil {
		return nil, err
	}

	cfg.Http.ReadTimeout, err = parseDuration("HTTP_READ_TIMEOUT", DefaultHttpReadTimeout)
	if err != nil {
		return nil, err
	}

	cfg.Http.WriteTimeout, err = parseDuration("HTTP_WRITE_TIMEOUT", DefaultHttpWriteTimeout)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
