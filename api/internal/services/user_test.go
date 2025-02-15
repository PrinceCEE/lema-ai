package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/princecee/lema-ai/config"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserServiceTestSuite struct {
	suite.Suite
	db          *gorm.DB
	userService *services.UserService
}

func (s *UserServiceTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	cfg.DSN = "file::memory:?cache=shared"
	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	s.userService = services.NewUserService(userRepo)

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
}

func (s *UserServiceTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
}

func (s *UserServiceTestSuite) TestUserService() {
	t := s.T()
	var users []*models.User

	t.Run("Get users", func(t *testing.T) {
		response, err := s.userService.GetUsers(1, 20)

		s.NoError(err)
		s.Equal(20, len(response.Users))

		users = response.Users
	})

	t.Run("Get user by ID", func(t *testing.T) {
		user, err := s.userService.GetUser(users[0].ID)

		s.NoError(err)
		s.NotNil(user)
		s.Equal(users[0].ID, user.ID)
	})

	t.Run("Get non-existing user", func(t *testing.T) {
		user, err := s.userService.GetUser(uuid.NewString())

		s.Error(err)
		s.Nil(user)
	})

	t.Run("Get user count", func(t *testing.T) {
		count, err := s.userService.GetUserCount()
		s.NoError(err)
		s.Equal(count, int64(20))
	})
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
