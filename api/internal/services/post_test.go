package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/princecee/lema-ai/config"
	database "github.com/princecee/lema-ai/internal/db"
	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/internal/db/repositories"
	"github.com/princecee/lema-ai/internal/services"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type PostServiceTestSuite struct {
	suite.Suite
	db          *gorm.DB
	postService *services.PostService
	users       []models.User
}

func (s *PostServiceTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	userRepo := repositories.NewUserRepository(db)
	postRepo := repositories.NewPostRepository(db)
	s.postService = services.NewPostService(postRepo)

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

		s.users = append(s.users, *user)
	}
}

func (s *PostServiceTestSuite) TearDownSuite() {
	sqlDB, err := s.db.DB()
	if err != nil {
		s.Fail(err.Error())
	}

	sqlDB.Close()
}

func (s *PostServiceTestSuite) TestCreatePost() {
	for _, user := range s.users {
		for i := 0; i < 5; i++ {
			now := time.Now()
			post := models.Post{
				Title:  gofakeit.Sentence(7),
				Body:   gofakeit.Sentence(40),
				UserID: user.ID,
			}

			err := s.postService.CreatePost(&post)
			s.NoError(err)
			s.NotEmpty(post.ID)
			s.NotEmpty(post.CreatedAt)
			s.NotEmpty(post.UpdatedAt)
			s.Equal(post.CreatedAt.After(now), true)
		}
	}
}

func (s *PostServiceTestSuite) TestGetPost() {
	for i := 1; i <= 10; i++ {
		post, err := s.postService.GetPost(uint(i))
		s.NoError(err)
		s.NotEmpty(post)
		s.Equal(post.ID, uint(i))
	}
}

func (s *PostServiceTestSuite) TestGetPosts() {
	for _, user := range s.users {
		posts, err := s.postService.GetPosts(user.ID)

		s.NoError(err)
		s.NotEmpty(posts)
		s.GreaterOrEqual(len(posts), 4)
	}
}

func (s *PostServiceTestSuite) TestDeletePost() {
	err := s.postService.DeletePost(20)
	s.NoError(err)

	post, err := s.postService.GetPost(20)
	s.Error(err)
	s.Nil(post)
}

func TestPostService(t *testing.T) {
	suite.Run(t, new(PostServiceTestSuite))
}
