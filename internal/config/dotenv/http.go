package dotenv

import (
	"os"

	"github.com/alirezaarzehgar/pastebin/internal/config"
)

const (
	DefaultHttpAddress      = "0.0.0.0"
	DefaultHttpPort         = 8080
	DefaultHttpIdleTimeout  = 120
	DefaultHttpReadTimeout  = 120
	DefaultHttpWriteTimeout = 120
)

func loadHttpConfig(cfg *config.Config) error {
	addr := os.Getenv("HTTP_ADDRESS")
	if addr == "" {
		addr = DefaultHttpAddress
	}
	cfg.Http.Address = addr

	var err error
	cfg.Http.Port, err = parseInt("HTTP_PORT", DefaultHttpPort)
	if err != nil {
		return err
	}

	cfg.Http.IdleTimeout, err = parseDuration("HTTP_IDLE_TIMEOUT", DefaultHttpIdleTimeout)
	if err != nil {
		return err
	}

	cfg.Http.ReadTimeout, err = parseDuration("HTTP_READ_TIMEOUT", DefaultHttpReadTimeout)
	if err != nil {
		return err
	}

	cfg.Http.WriteTimeout, err = parseDuration("HTTP_WRITE_TIMEOUT", DefaultHttpWriteTimeout)
	if err != nil {
		return err
	}

	return nil
}
