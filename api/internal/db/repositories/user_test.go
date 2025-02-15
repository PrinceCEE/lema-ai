package repositories_test

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
	"github.com/princecee/lema-ai/pkg/pagination"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db       *gorm.DB
	userRepo *repositories.UserRepository
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	cfg.DSN = "file::memory:?cache=shared"
	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	s.userRepo = repositories.NewUserRepository(db)
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
}

func (s *UserRepositoryTestSuite) TestUserRepository() {
	t := s.T()
	var userId string

	t.Run("Create user", func(t *testing.T) {
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

		err := s.userRepo.CreateUser(ctx, user)
		s.NoError(err)
		s.NotEmpty(user.ID)
		s.Equal(user.ID, user.Address.UserID)

		userId = user.ID
		_users := make([]*models.User, 19)
		for i := 0; i < 19; i++ {
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

			_users[i] = user
		}

		result := s.db.Create(&_users)
		s.NoError(result.Error)
	})

	t.Run("Get user count", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		count, err := s.userRepo.GetUserCount(ctx)
		s.NoError(err)
		s.Equal(count, int64(20))
	})

	t.Run("Get users", func(t *testing.T) {
		for i := 1; i <= 4; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			page := i
			limit := 5

			response, err := s.userRepo.GetUsers(ctx, pagination.PaginationQuery{
				Page:  &page,
				Limit: &limit,
			})

			s.NoError(err)
			s.Equal(limit, len(response.Users))
		}
	})

	t.Run("Get user by ID", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user, err := s.userRepo.GetUser(ctx, userId)
		s.NoError(err)
		s.NotEmpty(user)
		s.Equal(userId, user.ID)
	})
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
