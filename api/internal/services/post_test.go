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

type PostServiceTestSuite struct {
	suite.Suite
	db          *gorm.DB
	postService *services.PostService
	users       []models.User
}

func (s *PostServiceTestSuite) SetupSuite() {
	cfg := config.NewConfig("test", "silent")
	cfg.DSN = "file::memory:?cache=shared"
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

func (s *PostServiceTestSuite) TestPostService() {
	t := s.T()
	var postId string

	t.Run("Create post", func(t *testing.T) {
		for _, user := range s.users {
			for i := 0; i < 5; i++ {
				post := models.Post{
					ID:     uuid.NewString(),
					Title:  gofakeit.Sentence(7),
					Body:   gofakeit.Sentence(40),
					UserID: user.ID,
				}

				err := s.postService.CreatePost(&post)
				s.NoError(err)
				s.NotEmpty(post.ID)
				s.NotEmpty(post.CreatedAt)
			}
		}
	})

	t.Run("Get posts", func(t *testing.T) {
		posts, err := s.postService.GetPosts(s.users[0].ID)

		s.NoError(err)
		s.NotEmpty(posts)
		s.GreaterOrEqual(len(posts), 4)

		postId = posts[0].ID
	})

	t.Run("Get post by ID", func(t *testing.T) {
		post, err := s.postService.GetPost(postId)
		s.NoError(err)
		s.NotEmpty(post)
		s.Equal(postId, post.ID)
	})

	t.Run("Delete post", func(t *testing.T) {
		err := s.postService.DeletePost(postId)
		s.NoError(err)

		post, err := s.postService.GetPost(postId)
		s.Error(err)
		s.Nil(post)
	})
}

func TestPostService(t *testing.T) {
	suite.Run(t, new(PostServiceTestSuite))
}
