package http

import (
	"net/http"
)

func nweRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Worlds"))
	}))

	return mux
}
