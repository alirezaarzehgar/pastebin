package http

import (
	"fmt"
	"net/http"

	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/handler"
	"github.com/alirezaarzehgar/pastebin/internal/usecase"
)

type HttpHandler struct {
	config          *config.Config
	pastebinUsecase usecase.Pastebin
}

func New(
	cfg *config.Config,
	pastebinUsecase usecase.Pastebin,
) handler.Handler {
	return &HttpHandler{
		config:          cfg,
		pastebinUsecase: pastebinUsecase,
	}
}

func (h HttpHandler) Start() error {
	mux := h.newRouter()
	addr := fmt.Sprintf("%s:%d", h.config.Http.Address, h.config.Http.Port)

	srv := &http.Server{
		Addr:         addr,
		IdleTimeout:  h.config.Http.IdleTimeout,
		ReadTimeout:  h.config.Http.ReadTimeout,
		WriteTimeout: h.config.Http.WriteTimeout,
		Handler:      mux,
	}

	return srv.ListenAndServe()
}
