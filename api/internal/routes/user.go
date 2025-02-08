package routes

import (
	"github.com/go-chi/chi"
	"github.com/princecee/lema-ai/internal/handlers"
	"gorm.io/gorm"
)

func AddUserRoutes(db *gorm.DB, userService handlers.UserService) chi.Router {
	r := chi.NewRouter()
	h := handlers.NewUserHandler(userService)

	r.Get("/", h.GetUsers)
	r.Get("/count", h.GetUsersCount)
	r.Get("/{user_id}", h.GetUser)

	return r
}
