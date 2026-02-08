package main

import (
	"os"

	"github.com/alirezaarzehgar/pastebin/internal/logger"
	"github.com/alirezaarzehgar/pastebin/internal/usecase/pastebin"

	configFactory "github.com/alirezaarzehgar/pastebin/internal/config/factory"
	handlerHttp "github.com/alirezaarzehgar/pastebin/internal/handler/http"
	objStoreFactory "github.com/alirezaarzehgar/pastebin/internal/infra/objstorage/factory"
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

	objStore := objStoreFactory.New(objStoreFactory.ObjectStorageMinIO, cfg)

	err = objStore.Connect()
	if err != nil {
		logger.Error("unable to connect to object storage", "error", err)
		os.Exit(1)
	}
	logger.Debug("object storage is connected")

	pastebinUsecase := pastebin.New(nil, objStore, nil)

	logger.Info("start application", "address", cfg.Http.Address, "port", cfg.Http.Port)
	h := handlerHttp.New(cfg, pastebinUsecase)
	if err := h.Start(); err != nil {
		logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
