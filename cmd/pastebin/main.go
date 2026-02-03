package main

import (
	"os"

	configFactory "github.com/alirezaarzehgar/pastebin/internal/config/factory"
	handlerHttp "github.com/alirezaarzehgar/pastebin/internal/handler/http"
	"github.com/alirezaarzehgar/pastebin/internal/logger"
	loggerFactory "github.com/alirezaarzehgar/pastebin/internal/logger/factory"
)

func main() {
	configLoader := configFactory.New(configFactory.ConfigLoaderDotEnv)
	cfg, err := configLoader.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	loggerChoice := loggerFactory.New(loggerFactory.LoggerSlog, cfg)
	logger.DefaultLogger = loggerChoice

	logger.Info("start application", "address", cfg.Http.Address, "port", cfg.Http.Port)
	h := handlerHttp.New(cfg)
	if err := h.Start(); err != nil {
		logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
