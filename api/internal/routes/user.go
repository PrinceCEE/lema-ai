package routes

import (
	"github.com/go-chi/chi"
	"github.com/princecee/lema-ai/config"
	"github.com/princecee/lema-ai/internal/handlers"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func AddUserRoutes(db *gorm.DB, userService handlers.UserService, cfg *config.Config, l zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	h := handlers.NewUserHandler(userService, cfg, l)

	r.Get("/", h.GetUsers)
	r.Get("/count", h.GetUsersCount)
	r.Get("/{user_id}", h.GetUser)

	return r
}
