package repositories

import (
	"github.com/princecee/lema-ai/internal/db/models"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) CreatePost(p *models.Post) error {
	return r.db.Create(p).Error
}

func (r *PostRepository) GetPost(postId uint) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, postId).Error
	return &post, err
}

func (r *PostRepository) GetPosts(userId uint) ([]*models.Post, error) {
	var posts []*models.Post
	err := r.db.Where("user_id = ?", userId).Find(&posts).Error
	return posts, err
}

func (r *PostRepository) DeletePost(postId uint) error {
	result := r.db.Unscoped().Delete(&models.Post{}, postId)
	return result.Error
}
