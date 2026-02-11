package http

import (
	"net/http"
)

func (h *HttpHandler) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /paste", h.createPaste)
	mux.HandleFunc("GET /paste/{id}", h.getPasteContent)
	mux.HandleFunc("GET /paste/{id}/{filename}", h.getPasteFile)

	return mux
}
