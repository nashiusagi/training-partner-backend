package repositories

import (
	"github.com/jinzhu/gorm"
	"training-partner/domains"
)

type PostRepository interface {
	GetAll() ([]*domains.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetAll() ([]*domains.Post, error) {
	var posts []*domains.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
