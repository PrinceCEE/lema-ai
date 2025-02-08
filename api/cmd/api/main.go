package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/routes"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/princecee/lema-ai/pkg/response"
	"gorm.io/gorm"
)

func addRoutes(db *gorm.DB) chi.Router {
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	userRouter := routes.AddUserRoutes(db, userService)
	postRouter := routes.AddPostRoutes(db, postService)
	r := chi.NewRouter()

	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)
	r.Mount("/api/v1/users", userRouter)
	r.Mount("/api/v1/posts", postRouter)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := response.Response[any]{
			Message: fmt.Sprintf("%s %s not found", r.Method, r.URL.Path),
		}
		response.SendErrorResponse(w, resp, http.StatusNotFound)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := response.Response[any]{
			Message: fmt.Sprintf("%s %s not allowed", r.Method, r.URL.Path),
		}
		response.SendErrorResponse(w, resp, http.StatusMethodNotAllowed)
	})

	return r
}

func main() {
	db := database.GetDBConn()
	r := addRoutes(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", port),
	}

	errChan := make(chan error)
	log.Printf("Server started on port :%s", port)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case <-errChan:
	case <-ctx.Done():
		stop()
		log.Println("Server shutting down in 5s")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server shut down successfully")
}
