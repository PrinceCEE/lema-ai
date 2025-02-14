package services

import (
	"context"
	"errors"
	"time"

	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(ctx context.Context, p *models.Post) error
	GetPost(ctx context.Context, postId uint) (*models.Post, error)
	GetPosts(ctx context.Context, userId uint) ([]*models.Post, error)
	DeletePost(ctx context.Context, postId uint) error
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo}
}

func (s *PostService) CreatePost(p *models.Post) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.postRepo.CreatePost(ctx, p)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return apperror.ErrNotFound
		default:
			return apperror.ErrInternalServer
		}
	}

	return nil
}

func (s *PostService) GetPost(postId uint) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post, err := s.postRepo.GetPost(ctx, postId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, apperror.ErrNotFound
		default:
			return nil, apperror.ErrInternalServer
		}
	}

	return post, nil
}

func (s *PostService) GetPosts(userId uint) ([]*models.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	posts, err := s.postRepo.GetPosts(ctx, userId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return []*models.Post{}, nil
		default:
			return nil, apperror.ErrInternalServer
		}
	}

	return posts, nil
}

func (s *PostService) DeletePost(postId uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.postRepo.DeletePost(ctx, postId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.ErrInternalServer
	}

	return nil
}
