package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

type UserHandlerTestSuite struct {
	suite.Suite
	db     *gorm.DB
	server *httptest.Server
}

func (s *UserHandlerTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "debug")
	var logger zerolog.Logger

	db := database.GetDBConn(cfg.DSN)

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	for i := 0; i < 20; i++ {
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
		err := userRepo.CreateUser(user)
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

	t.Run("Get users with pagination query", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "?page=1&limit=5")
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[[]*models.User]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(true, *response.Success)
		s.Equal("Users fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.Equal(5, len(response.Data))
	})

	t.Run("Get users without pagination query", func(t *testing.T) {
		resp, err := s.server.Client().Get(url)
		s.NoError(err)
		s.Equal(http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()

		response := response.Response[[]*models.User]{}
		_ = json.ReadJSON(resp.Body, &response)

		s.Equal(true, *response.Success)
		s.Equal("Users fetched successfully", response.Message)
		s.NotEmpty(response.Data)
		s.Equal(10, len(response.Data))
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
		resp, err := s.server.Client().Get(url + "/1")
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

	t.Run("Get user by out of range ID", func(t *testing.T) {
		resp, err := s.server.Client().Get(url + "/100")
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
