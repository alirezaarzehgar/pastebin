package dotenv

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
)

const (
	DefaultHandlerMaxFileSize int64 = (1 << 10) * 10
)

func loadHandlerConfig(cfg *config.Config) error {
	maxFileSize, err := parseInt64(KeyHandlerMaxFileSize, DefaultHandlerMaxFileSize)
	if err != nil {
		return err
	}
	cfg.Handler.MaxFileSize = maxFileSize

	return nil
}
