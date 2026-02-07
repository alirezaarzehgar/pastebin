package http

import (
	"net/http"
)

func (h *HttpHandler) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/paste", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.createPaste(w, r)
		case http.MethodGet:
			h.getPaste(w, r)
		}
	})

	return mux
}
