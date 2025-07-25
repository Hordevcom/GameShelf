package routes

import (
	"github.com/Hordevcom/GameShelf/internal/handlers"
	"github.com/Hordevcom/GameShelf/internal/middleware/auth"
	"github.com/Hordevcom/GameShelf/internal/middleware/logging"
	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler, log *logging.Logger) *chi.Mux {
	router := chi.NewRouter()

	router.Use(log.WithLogging)

	router.Post("/api/user/register", h.UserRegister())
	router.Post("/api/user/login", h.Login())
	router.With(auth.AuthMiddleware).Post("/api/games", h.AddNewGame())

	return router
}
