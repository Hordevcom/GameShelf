package routes

import (
	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler, log *logging.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(log.WithLogging)

	return router
}
