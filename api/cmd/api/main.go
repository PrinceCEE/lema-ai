package main

import (
	"context"
	"flag"
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
	"github.com/joho/godotenv"
	"github.com/princecee/lema-ai/config"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/db/seeder"
	"github.com/princecee/lema-ai/internal/middlewares"
	"github.com/princecee/lema-ai/internal/routes"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/princecee/lema-ai/pkg/response"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func addRoutes(db *gorm.DB, cfg *config.Config, l zerolog.Logger) chi.Router {
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	userRouter := routes.AddUserRoutes(db, userService, cfg, l)
	postRouter := routes.AddPostRoutes(db, postService, cfg, l)
	r := chi.NewRouter()

	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(middleware.CleanPath)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)
	r.Use(middlewares.RequestSize(1 << 20)) // 1mb body limit
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

	if cfg.ENV == config.EnvDevelopment {
		seeder.Seed(userRepo, postRepo)
	}

	return r
}

func main() {
	var env, loglevel string

	flag.StringVar(&env, "env", config.EnvDevelopment, "Environment to run the server")
	flag.StringVar(&loglevel, "loglevel", "silent", "Log level for the server")
	flag.Parse()

	if env != "test" {
		_ = godotenv.Load()
	}

	cfg := config.NewConfig(env, loglevel)
	logger := zerolog.New(os.Stdout).Level(config.GetLoggerLevel(cfg.LOG_LEVEL))

	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)
	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		panic(err)
	}
	r := addRoutes(db, cfg, logger)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := http.Server{
		Handler: r,
		Addr:    fmt.Sprintf(":%s", cfg.PORT),
	}

	errChan := make(chan error)
	log.Printf("Server started on port :%s", cfg.PORT)
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
