package services

import (
	"errors"

	"github.com/princecee/lema-ai/internal/db/models"
	apperror "github.com/princecee/lema-ai/pkg/error"
	"github.com/princecee/lema-ai/pkg/pagination"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(userId uint) (*models.User, error)
	GetUsers(opts pagination.PaginationQuery) ([]*models.User, error)
	GetUserCount() (int64, error)
}

type UserService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetUser(userId uint) (*models.User, error) {
	user, err := s.userRepo.GetUser(userId)
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

func (s *UserService) GetUsers(page, limit int) ([]*models.User, error) {
	users, err := s.userRepo.GetUsers(pagination.PaginationQuery{
		Page:  &page,
		Limit: &limit,
	})

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return []*models.User{}, nil
		default:
			return nil, apperror.ErrInternalServer
		}
	}

	return users, nil
}

func (s *UserService) GetUserCount() (int64, error) {
	count, err := s.userRepo.GetUserCount()
	if err != nil {
		return 0, apperror.ErrInternalServer
	}

	return count, nil
}
