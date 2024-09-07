package repositories

import (
	"training-partner/internal/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:generate mockery --name ExerciseRepository
type ExerciseRepository interface {
	GetAll() ([]*domains.Exercise, error)
	FindById(id int) (*domains.Exercise, error)
}

type exerciseRepository struct {
	db *gorm.DB
}

func NewExerciseRepository(db *gorm.DB) ExerciseRepository {
	return &exerciseRepository{db}
}

func (r *exerciseRepository) GetAll() ([]*domains.Exercise, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var exercises []*domains.Exercise
	if err := r.db.Preload("Muscles").Find(&exercises).Error; err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *exerciseRepository) FindById(id int) (*domains.Exercise, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var exercise *domains.Exercise
	if err := r.db.Preload("Muscles").Find(&exercise, id).Error; err != nil {
		return nil, err
	}
	return exercise, nil
}
