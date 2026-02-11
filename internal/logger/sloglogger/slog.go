package sloglogger

import (
	"log/slog"
	"os"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
)

type slogLogger struct {
	logger *slog.Logger
}

func (l slogLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l slogLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l slogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l slogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l slogLogger) Verbose(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

type Options func(l *slogLogger)

func OptionSetLogLevel(level logger.Leveler) Options {
	var slogLevel slog.Level
	switch level {
	case logger.LevelError:
		slogLevel = slog.LevelError
	case logger.LevelWarn:
		slogLevel = slog.LevelWarn
	case logger.LevelInfo:
		slogLevel = slog.LevelInfo
	case logger.LevelDebug:
		slogLevel = slog.LevelDebug
	case logger.LevelVerbose:
		slogLevel = slog.LevelDebug - 2
	default:
		slogLevel = slog.LevelInfo
	}

	return func(l *slogLogger) {
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slogLevel,
		})
	}
}

func New(opts ...Options) logger.Logger {
	l := &slogLogger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}
