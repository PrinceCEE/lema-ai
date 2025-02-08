package repositories_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PostRepositoryTestSuite struct {
	suite.Suite
	db       *gorm.DB
	userRepo *repositories.UserRepository
	postRepo *repositories.PostRepository
	users    []models.User
}

func (s *PostRepositoryTestSuite) SetupSuite() {
	db := database.GetDBConn()

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	s.userRepo = repositories.NewUserRepository(db)
	s.postRepo = repositories.NewPostRepository(db)

	for i := 0; i < 5; i++ {
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
		if err != nil {
			s.Fail(err.Error())
		}

		s.users = append(s.users, *user)
	}
}

func (s *PostRepositoryTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
}

func (s *PostRepositoryTestSuite) TestCreatePost() {
	for _, user := range s.users {
		for i := 0; i < 5; i++ {
			now := time.Now()
			post := models.Post{
				Title:  gofakeit.Sentence(10),
				Body:   gofakeit.Sentence(1),
				UserID: user.ID,
			}

			err := s.postRepo.CreatePost(&post)
			s.NoError(err)
			s.NotEmpty(post.ID)
			s.NotEmpty(post.CreatedAt)
			s.NotEmpty(post.UpdatedAt)
			s.Equal(post.CreatedAt.After(now), true)
		}
	}
}

func (s *PostRepositoryTestSuite) TestGetPost() {
	for i := 1; i <= 10; i++ {
		post, err := s.postRepo.GetPost(uint(i))
		s.NoError(err)
		s.NotEmpty(post)
		s.Equal(post.ID, uint(i))
	}
}

func (s *PostRepositoryTestSuite) TestGetPosts() {
	for _, user := range s.users {
		posts, err := s.postRepo.GetPosts(user.ID)

		s.NoError(err)
		s.NotEmpty(posts)
		s.GreaterOrEqual(len(posts), 4)
	}
}

func (s *PostRepositoryTestSuite) TestDeletePost() {
	err := s.postRepo.DeletePost(20)
	s.NoError(err)

	post, err := s.postRepo.GetPost(20)
	s.Error(err)
	s.NotEqual(post.ID, 20)
}

func TestPostRepository(t *testing.T) {
	suite.Run(t, new(PostRepositoryTestSuite))
}
