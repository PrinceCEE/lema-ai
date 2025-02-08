package services_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
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
	db := database.GetDBConn()

	err := db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	s.userService = services.NewUserService(userRepo)

	for i := 0; i < 20; i++ {
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
		err := userRepo.CreateUser(user)
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

func (s *UserServiceTestSuite) TestGetUser() {
	t := s.T()

	t.Run("Get existing users", func(t *testing.T) {
		for i := 1; i <= 20; i++ {
			user, err := s.userService.GetUser(uint(i))

			s.NoError(err)
			s.NotNil(user)
			s.Equal(user.ID, uint(i))
		}
	})

	t.Run("Get non-existing user", func(t *testing.T) {
		user, err := s.userService.GetUser(21)

		s.Error(err)
		s.Nil(user)
	})
}

func (s *UserServiceTestSuite) TestGetUsers() {
	for i := 1; i <= 4; i++ {
		users, err := s.userService.GetUsers(i, 5)

		s.NoError(err)
		s.Equal(len(users), 5)
	}
}

func (s *UserServiceTestSuite) TestGetUserCount() {
	count, err := s.userService.GetUserCount()
	s.NoError(err)
	s.Equal(count, int64(20))
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
