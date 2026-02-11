package http

import (
	"net/http"
)

func (h *HttpHandler) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /paste", h.createPaste)
	mux.HandleFunc("GET /paste/{id}", h.getPaste)

	return mux
}
