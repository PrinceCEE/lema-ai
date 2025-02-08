package repositories

import (
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

func (r *UserRepository) CreateUser(u *models.User) error {
	result := r.db.Create(u)
	return result.Error
}

func (r *UserRepository) GetUser(userId uint) (*models.User, error) {
	u := models.User{}
	err := r.db.First(&u, userId).Error
	return &u, err
}

func (r *UserRepository) GetUsers(opts pagination.PaginationQuery) ([]*models.User, error) {
	offset := pagination.GetPaginationData(opts)
	var users []*models.User
	err := r.db.Offset(offset).Limit(*opts.Limit).Find(&users).Error

	return users, err
}

func (r *UserRepository) GetUserCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	return count, err
}
