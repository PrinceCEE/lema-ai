package repositories

import (
	"context"

	"github.com/princecee/lema-ai/internal/db/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) CreatePost(ctx context.Context, p *models.Post) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *PostRepository) GetPost(ctx context.Context, postId string) (*models.Post, error) {
	var post models.Post
	err := r.db.WithContext(ctx).Where("id = ?", postId).First(&post).Error
	return &post, err
}

func (r *PostRepository) GetPosts(ctx context.Context, userId string) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) DeletePost(ctx context.Context, postId string) error {
	result := r.db.WithContext(ctx).Unscoped().Where("id = ?", postId).Delete(&models.Post{})
	return result.Error
}
