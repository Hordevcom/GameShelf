package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer(router *chi.Mux) http.Server {
	return http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
}
