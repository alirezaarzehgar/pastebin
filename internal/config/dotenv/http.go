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
	addr := os.Getenv(KeyHttpAddress)
	if addr == "" {
		addr = DefaultHttpAddress
	}
	cfg.Http.Address = addr

	var err error
	cfg.Http.Port, err = parseInt(KeyHttpPort, DefaultHttpPort)
	if err != nil {
		return err
	}

	cfg.Http.IdleTimeout, err = parseDuration(KeyHttpIdleTimeout, DefaultHttpIdleTimeout)
	if err != nil {
		return err
	}

	cfg.Http.ReadTimeout, err = parseDuration(KeyHttpReadTimeout, DefaultHttpReadTimeout)
	if err != nil {
		return err
	}

	cfg.Http.WriteTimeout, err = parseDuration(KeyHttpWriteTimeout, DefaultHttpWriteTimeout)
	if err != nil {
		return err
	}

	return nil
}
