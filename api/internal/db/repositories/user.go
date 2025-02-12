package repositories

import (
	"context"

	"github.com/princecee/lema-ai/internal/db/models"
	"github.com/princecee/lema-ai/pkg/pagination"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(ctx context.Context, u *models.User) error {
	result := r.db.WithContext(ctx).Create(u)
	return result.Error
}

func (r *UserRepository) GetUser(ctx context.Context, userId uint) (*models.User, error) {
	u := models.User{}
	err := r.db.WithContext(ctx).Preload("Address").First(&u, userId).Error
	return &u, err
}

func (r *UserRepository) GetUsers(ctx context.Context, opts pagination.PaginationQuery) (*pagination.GetUsersResult, error) {
	count, err := r.GetUserCount(ctx)
	if err != nil {
		return nil, err
	}

	offset := pagination.GetPaginationData(opts)
	var users []*models.User
	err = r.db.WithContext(ctx).Preload("Address").Offset(offset).Limit(*opts.Limit).Find(&users).Error

	totalPages := pagination.GetTotalPages(count, *opts.Limit)
	return &pagination.GetUsersResult{
		Users:      users,
		Count:      count,
		TotalPages: totalPages,
		Page:       int64(*opts.Page),
		Limit:      int64(*opts.Limit),
		HasNext:    *opts.Page < int(totalPages),
		HasPrev:    *opts.Page > 1,
	}, err
}

func (r *UserRepository) GetUserCount(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.User{}).Count(&count).Error
	return count, err
}
