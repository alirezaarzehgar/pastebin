package dotenv

import (
	"fmt"
	"os"
	"strings"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
)

func loadLogLevel() (logger.Leveler, error) {
	ll := os.Getenv(KeyLogLevel)
	ll = strings.ToLower(ll)

	switch ll {
	case "error", "err":
		return logger.LevelError, nil
	case "warning", "warn":
		return logger.LevelWarn, nil
	case "information", "info":
		return logger.LevelInfo, nil
	case "debug":
		return logger.LevelDebug, nil
	case "verbose":
		return logger.LevelVerbose, nil
	case "":
		return logger.LevelInfo, nil
	default:
		return logger.LevelUnknown, fmt.Errorf("invalid log level: %s", ll)
	}
}
