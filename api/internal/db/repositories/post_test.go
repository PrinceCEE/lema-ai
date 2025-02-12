package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/princecee/lema-ai/config"
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
	cfg := config.NewConfig("test", "silent")
	db := database.GetDBConn(cfg.DSN, cfg.MAX_IDLE_CONNS, cfg.MAX_OPEN_CONNS, cfg.CONN_MAX_LIFETIME, cfg.LOG_LEVEL)

	err := db.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{})
	if err != nil {
		s.Fail(err.Error())
	}

	s.db = db
	s.userRepo = repositories.NewUserRepository(db)
	s.postRepo = repositories.NewPostRepository(db)

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
		err := s.userRepo.CreateUser(ctx, user)
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

func (s *PostRepositoryTestSuite) TestPostRepository() {
	t := s.T()

	t.Run("Create post", func(t *testing.T) {
		for _, user := range s.users {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			for i := 0; i < 5; i++ {
				now := time.Now()
				post := models.Post{
					Title:  gofakeit.Sentence(7),
					Body:   gofakeit.Sentence(40),
					UserID: user.ID,
				}

				err := s.postRepo.CreatePost(ctx, &post)
				s.NoError(err)
				s.NotEmpty(post.ID)
				s.NotEmpty(post.CreatedAt)
				s.NotEmpty(post.UpdatedAt)
				s.Equal(post.CreatedAt.After(now), true)
			}
		}
	})

	t.Run("Get post", func(t *testing.T) {
		for _, user := range s.users {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			posts, err := s.postRepo.GetPosts(ctx, user.ID)

			s.NoError(err)
			s.NotEmpty(posts)
			s.GreaterOrEqual(len(posts), 4)
		}
	})

	t.Run("Get post by ID", func(t *testing.T) {
		for i := 1; i <= 10; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			post, err := s.postRepo.GetPost(ctx, uint(i))
			s.NoError(err)
			s.NotEmpty(post)
			s.Equal(post.ID, uint(i))
		}
	})

	t.Run("Delete post", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := s.postRepo.DeletePost(ctx, 20)
		s.NoError(err)

		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		post, err := s.postRepo.GetPost(ctx, 20)
		s.Error(err)
		s.NotEqual(post.ID, 20)
	})
}

func TestPostRepository(t *testing.T) {
	suite.Run(t, new(PostRepositoryTestSuite))
}
