package services

import (
	"context"
	"errors"
	"time"

	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"github.com/princecee/lema-ai/pkg/pagination"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(ctx context.Context, userId uint) (*models.User, error)
	GetUsers(ctx context.Context, opts pagination.PaginationQuery) (*pagination.GetUsersResult, error)
	GetUserCount(ctx context.Context) (int64, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetUser(userId uint) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := s.userRepo.GetUser(ctx, userId)
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, apperror.ErrNotFound
		default:
			return nil, apperror.ErrInternalServer
		}
	}

	return user, nil
}

func (s *UserService) GetUsers(page, limit int) (*pagination.GetUsersResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	users, err := s.userRepo.GetUsers(ctx, pagination.PaginationQuery{
		Page:  &page,
		Limit: &limit,
	})

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return &pagination.GetUsersResult{
				Users: []*models.User{},
			}, nil
		default:
			return nil, apperror.ErrInternalServer
		}
	}

	return users, nil
}

func (s *UserService) GetUserCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := s.userRepo.GetUserCount(ctx)
	if err != nil {
		return 0, apperror.ErrInternalServer
	}

	return count, nil
}
