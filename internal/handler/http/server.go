package http

import (
	"fmt"
	"net/http"

	"github.com/alirezaarzehgar/pastebin/internal/config"
	"github.com/alirezaarzehgar/pastebin/internal/handler"
)

type HttpHandler struct {
	config *config.Config
}

func New(cfg *config.Config) handler.Handler {
	return &HttpHandler{
		config: cfg,
	}
}

func (h HttpHandler) Start() error {
	mux := nweRouter()
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
