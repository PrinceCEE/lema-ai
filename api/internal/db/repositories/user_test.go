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
	db := database.GetDBConn()

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

func (s *UserRepositoryTestSuite) TestCreateUser() {
	for i := 0; i < 20; i++ {
		now := time.Now()

		user := &models.User{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			Username:  gofakeit.Username(),
			Phone:     gofakeit.Phone(),
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
	}
}

func (s *UserRepositoryTestSuite) TestGetUser() {
	for i := 1; i <= 20; i++ {
		user, err := s.userRepo.GetUser(uint(i))
		s.NoError(err)
		s.NotEmpty(user)
		s.Equal(user.ID, uint(i))
	}
}

func (s *UserRepositoryTestSuite) TestGetUsers() {
	for i := 1; i <= 4; i++ {
		page := i
		limit := 5

		users, err := s.userRepo.GetUsers(pagination.PaginationQuery{
			Page:  &page,
			Limit: &limit,
		})

		s.NoError(err)
		s.Equal(len(users), limit)
	}
}

func (s *UserRepositoryTestSuite) TestGetUserCount() {
	count, err := s.userRepo.GetUserCount()
	s.NoError(err)
	s.Equal(count, int64(20))
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
