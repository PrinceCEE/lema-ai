package handlers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/princecee/lema-ai/config"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/routes"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/princecee/lema-ai/pkg/json"
	"github.com/princecee/lema-ai/pkg/pagination"
	"github.com/princecee/lema-ai/pkg/response"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserHandlerTestSuite struct {
	suite.Suite
	db     *gorm.DB
	server *httptest.Server
}

func (s *UserHandlerTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	cfg.DSN = "file::memory:?cache=shared"
	var logger zerolog.Logger

	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	for i := 0; i < 20; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user := &models.User{
			ID:       uuid.NewString(),
			Name:     gofakeit.Name(),
			Username: gofakeit.Username(),
			Phone:    gofakeit.Phone(),
			Email:    gofakeit.Email(),
			Address: models.Address{
				ID:      uuid.NewString(),
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
	userRouter := routes.AddUserRoutes(db, userService, cfg, logger)
	r.Mount("/api/v1/users", userRouter)

	s.server = httptest.NewServer(r)
}

func (s *UserHandlerTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
	s.server.Close()
}

func (s *UserHandlerTestSuite) TestUserHandler() {
	t := s.T()
	url := s.server.URL + "/api/v1/users"
	var userId string

	t.Run("Get users with pagination query", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "?page=1&limit=5")
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[*pagination.GetUsersResult]{}
		_ = json.ReadJSON(resp.Body, &response)

		userLen := len(response.Data.Users)
		s.Equal(true, *response.Success)
		s.Equal("Users fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.Equal(int64(5), int64(userLen))
		s.Equal(int64(1), response.Data.Page)
		s.Equal(int64(5), response.Data.Limit)

		userId = response.Data.Users[0].ID
	})

	t.Run("Get users without pagination query", func(t *testing.T) {
		resp, err := s.server.Client().Get(url)
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[*pagination.GetUsersResult]{}
		_ = json.ReadJSON(resp.Body, &response)

		userLen := len(response.Data.Users)
		s.Equal(true, *response.Success)
		s.Equal("Users fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.Equal(int64(10), int64(userLen))
		s.Equal(int64(1), response.Data.Page)
		s.Equal(int64(10), response.Data.Limit)
	})

	t.Run("Get users count", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/count")
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[map[string]int64]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(true, *response.Success)
		s.Equal("Users count fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.Equal(int64(20), response.Data["count"])
	})

	t.Run("Get user by ID", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/" + userId)
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[models.User]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(true, *response.Success)
		s.Equal("User fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.NotEmpty(response.Data.ID)
	})

	t.Run("Get user with non-existent ID", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/" + uuid.NewString())
		s.NoError(err)
		s.Equal(http.StatusNotFound, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(false, *response.Success)
		s.Equal("not found", response.Message)
		s.Empty(response.Data)
	})

	t.Run("Get user by invalid ID", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/invalid")
		s.NoError(err)
		s.Equal(http.StatusBadRequest, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[any]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(false, *response.Success)
		s.Equal("Invalid user ID", response.Message)
		s.Empty(response.Data)
	})
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
