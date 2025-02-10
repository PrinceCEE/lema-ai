package routes

import (
	"github.com/go-chi/chi"
	"github.com/princecee/lema-ai/config"
	"github.com/princecee/lema-ai/internal/handlers"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func AddPostRoutes(db *gorm.DB, postService handlers.PostService, cfg *config.Config, l zerolog.Logger) chi.Router {
	r := chi.NewRouter()
	h := handlers.NewPostHandler(postService, cfg, l)

	r.Post("/", h.CreatePost)
	r.Get("/", h.GetPosts)
	r.Get("/{post_id}", h.GetPost)
	r.Delete("/{post_id}", h.DeletePost)

	return r
}
