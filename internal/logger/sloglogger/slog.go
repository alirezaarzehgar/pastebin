package sloglogger

import (
	"log/slog"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
)

type slogLogger struct {
	level logger.Leveler
}

func (l slogLogger) Error(msg string, args ...any) {
	slog.Error(msg, args...)
}

func (l slogLogger) Warn(msg string, args ...any) {
	if l.level > logger.LevelWarn {
		return
	}
	slog.Warn(msg, args...)

}

func (l slogLogger) Info(msg string, args ...any) {
	if l.level > logger.LevelInfo {
		return
	}
	slog.Info(msg, args...)
}

func (l slogLogger) Debug(msg string, args ...any) {
	if l.level > logger.LevelDebug {
		return
	}
	slog.Debug(msg, args...)
}

func (l slogLogger) Verbose(msg string, args ...any) {
	if l.level > logger.LevelVerbose {
		return
	}
	slog.Debug(msg, args...)
}

type Options func(l *slogLogger)

func OptionSetLogLevel(level logger.Leveler) Options {
	return func(l *slogLogger) {
		l.level = level
	}
}

func New(opts ...Options) logger.Logger {
	l := &slogLogger{
		level: logger.LevelInfo,
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}
