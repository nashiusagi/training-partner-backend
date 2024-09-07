package repositories

import (
	"training-partner/internal/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MuscleRepository interface {
	GetAll() ([]*domains.Muscle, error)
	FindById(id int) (*domains.Muscle, error)
}

type muscleRepository struct {
	db *gorm.DB
}

func NewMuscleRepository(db *gorm.DB) MuscleRepository {
	return &muscleRepository{db}
}

func (r *muscleRepository) GetAll() ([]*domains.Muscle, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var muscles []*domains.Muscle
	if err := r.db.Find(&muscles).Error; err != nil {
		return nil, err
	}
	return muscles, nil
}

func (r *muscleRepository) FindById(id int) (*domains.Muscle, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var muscle *domains.Muscle
	if err := r.db.Find(&muscle, id).Error; err != nil {
		return nil, err
	}
	return muscle, nil
}
