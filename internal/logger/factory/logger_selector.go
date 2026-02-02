package factory

import (
	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/logger"
	"github.com/alirezaarzehgar/pastebin/internal/logger/sloglogger"
)

type Logger uint

const (
	LoggerSlog = iota
)

func New(loggerChoice Logger, cfg *config.Config) logger.Logger {
	switch loggerChoice {
	case LoggerSlog:
		return sloglogger.New(
			sloglogger.OptionSetLogLevel(cfg.LogLevel),
		)
	default:
		panic("invalid Logger choice")
	}
}
