package repositories

import (
	"training-partner/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostRepository interface {
	GetAll() ([]*domains.Post, error)
	FindById(id int) (*domains.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) GetAll() ([]*domains.Post, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var posts []*domains.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *postRepository) FindById(id int) (*domains.Post, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var post *domains.Post
	if err := r.db.Find(&post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}
