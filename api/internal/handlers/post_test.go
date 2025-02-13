package handlers_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi"
	"github.com/princecee/lema-ai/config"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/routes"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/princecee/lema-ai/pkg/json"
	"github.com/princecee/lema-ai/pkg/response"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PostHandlerTestSuite struct {
	suite.Suite
	db     *gorm.DB
	server *httptest.Server
}

func (s *PostHandlerTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	var logger zerolog.Logger

	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	for i := 0; i < 5; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user := &models.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Phone:     gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Address: models.Address{
				Street:  gofakeit.StreetName(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				Zipcode: gofakeit.Zip(),
			},
		}
		err := userRepo.CreateUser(ctx, user)
		if err != nil {
			s.Fail(err.Error())
		}
	}

	r := chi.NewRouter()
	postRouter := routes.AddPostRoutes(db, postService, cfg, logger)
	r.Mount("/api/v1/posts", postRouter)

	s.server = httptest.NewServer(r)
}

func (s *PostHandlerTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
	s.server.Close()
}

func (s *PostHandlerTestSuite) TestPostHandler() {
	t := s.T()
	url := s.server.URL + "/api/v1/posts"

	t.Run("Create new posts", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			for j := 1; j <= 5; j++ {
				payload, _ := json.WriteJSON(map[string]any{
					"title":  gofakeit.Sentence(7),
					"body":   gofakeit.Sentence(40),
					"userId": i,
				})

				resp, err := s.server.Client().Post(url, "application/json", bytes.NewBuffer(payload))
				s.NoError(err)
				s.Equal(http.StatusOK, resp.StatusCode)
				defer resp.Body.Close()

				response := response.Response[*models.Post]{}
				_ = json.ReadJSON(resp.Body, &response)

				s.Equal(true, *response.Success)
				s.Equal("Post created successfully", response.Message)
				s.NotEmpty(response.Data)
				s.NotEmpty(response.Data.ID)
				s.Equal(uint(i), response.Data.UserID)
			}
		}
	})

	t.Run("Get post by ID", func(t *testing.T) {
		for i := 1; i <= 25; i++ {
			resp, err := s.server.Client().Get(url + fmt.Sprintf("/%d", i))
			s.NoError(err)
			s.Equal(http.StatusOK, resp.StatusCode)
			defer resp.Body.Close()

			response := response.Response[*models.Post]{}
			_ = json.ReadJSON(resp.Body, &response)
			s.Equal(true, *response.Success)
			s.Equal("Post fetched successfully", response.Message)
			s.NotEmpty(response.Data)
			s.Equal(uint(i), response.Data.ID)
		}
	})

	t.Run("Get non-existent post by ID", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/100")
		s.NoError(err)
		s.Equal(http.StatusNotFound, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)
		s.Equal(false, *response.Success)
		s.Equal("not found", response.Message)
	})

	t.Run("Get posts by user ID", func(t *testing.T) {
		for i := 1; i <= 5; i++ {
			resp, err := s.server.Client().Get(url + fmt.Sprintf("?user_id=%d", i))
			s.NoError(err)
			s.Equal(http.StatusOK, resp.StatusCode)
			defer resp.Body.Close()

			response := response.Response[[]*models.Post]{}
			_ = json.ReadJSON(resp.Body, &response)
			s.Equal(true, *response.Success)
			s.Equal("Posts fetched successfully", response.Message)
			s.NotEmpty(response.Data)
			s.Equal(len(response.Data), 5)
			s.NotEmpty(response.Data[0].ID)
		}
	})

	t.Run("Get posts with wrong user id", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "?user_id=c")
		s.NoError(err)
		s.Equal(http.StatusBadRequest, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)
		s.Equal(false, *response.Success)
		s.Equal("Invalid user ID", response.Message)
		s.Empty(response.Data)
	})

	t.Run("Delete post by ID", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, url+"/1", nil)
		s.NoError(err)

		resp, err := s.server.Client().Do(req)
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)
		s.Equal(true, *response.Success)
		s.Equal("Post deleted successfully", response.Message)
		s.Empty(response.Data)
	})

	t.Run("Fetch deleted post", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/1")
		s.NoError(err)
		s.Equal(http.StatusNotFound, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)
		s.Equal(false, *response.Success)
		s.Equal("not found", response.Message)
	})
}

func TestPostHandler(t *testing.T) {
	suite.Run(t, new(PostHandlerTestSuite))
}
