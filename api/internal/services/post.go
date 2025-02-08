package services

import (
	"errors"

	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(p *models.Post) error
	GetPost(postId uint) (*models.Post, error)
	GetPosts(userId uint) ([]*models.Post, error)
	DeletePost(postId uint) error
}

type PostService struct {
	postRepo PostRepository
}

func NewPostService(postRepo PostRepository) *PostService {
	return &PostService{postRepo}
}

func (s *PostService) CreatePost(p *models.Post) error {
	err := s.postRepo.CreatePost(p)
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
	post, err := s.postRepo.GetPost(postId)
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
	posts, err := s.postRepo.GetPosts(userId)
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
	err := s.postRepo.DeletePost(postId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return apperror.ErrInternalServer
	}

	return nil
}
