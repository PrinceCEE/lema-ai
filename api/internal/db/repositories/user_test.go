package repositories_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
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
	db := database.GetDBConn("file::memory:?cache=shared")

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

	t.Run("Create user", func(t *testing.T) {
		now := time.Now()
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

		err := s.userRepo.CreateUser(user)
		s.NoError(err)
		s.NotEmpty(user.ID)
		s.NotEmpty(user.CreatedAt)
		s.NotEmpty(user.UpdatedAt)
		s.Equal(user.CreatedAt.After(now), true)
		s.Equal(user.ID, user.Address.UserID)

		_users := make([]*models.User, 19)
		for i := 0; i < 19; i++ {
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

			_users[i] = user
		}

		result := s.db.Create(&_users)
		s.NoError(result.Error)
	})

	t.Run("Get user count", func(t *testing.T) {
		count, err := s.userRepo.GetUserCount()
		s.NoError(err)
		s.Equal(count, int64(20))
	})

	t.Run("Get users", func(t *testing.T) {
		for i := 1; i <= 4; i++ {
			page := i
			limit := 5

			response, err := s.userRepo.GetUsers(pagination.PaginationQuery{
				Page:  &page,
				Limit: &limit,
			})

			s.NoError(err)
			s.Equal(limit, len(response.Users))
		}
	})

	t.Run("Get user by ID", func(t *testing.T) {
		user, err := s.userRepo.GetUser(10)

		s.NoError(err)
		s.NotEmpty(user)
		s.Equal(uint(10), user.ID)
	})
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
