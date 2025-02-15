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
	cfg.DSN = "file::memory:?cache=shared"
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
				post := models.Post{
					ID:        uuid.NewString(),
					Title:     gofakeit.Sentence(7),
					Body:      gofakeit.Sentence(40),
					UserID:    user.ID,
					CreatedAt: time.Now().UTC().Format(time.RFC3339),
				}

				err := s.postRepo.CreatePost(ctx, &post)
				s.NoError(err)
				s.NotEmpty(post.ID)
				s.NotEmpty(post.CreatedAt)
			}
		}
	})

	t.Run("Get post", func(t *testing.T) {
		user := s.users[0]

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		posts, err := s.postRepo.GetPosts(ctx, user.ID)

		s.NoError(err)
		s.NotEmpty(posts)
		s.GreaterOrEqual(len(posts), 4)

		t.Run("Get post by ID", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			post, err := s.postRepo.GetPost(ctx, posts[0].ID)
			s.NoError(err)
			s.NotEmpty(post)
			s.Equal(posts[0].ID, post.ID)
		})

		t.Run("Delete post", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := s.postRepo.DeletePost(ctx, posts[0].ID)
			s.NoError(err)

			ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			post, err := s.postRepo.GetPost(ctx, posts[0].ID)
			s.Error(err)
			s.Empty(post)
		})
	})
}

func TestPostRepository(t *testing.T) {
	suite.Run(t, new(PostRepositoryTestSuite))
}
