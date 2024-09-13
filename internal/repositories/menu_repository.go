package repositories

import (
	"training-partner/internal/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:generate mockery --name MenuRepository
type MenuRepository interface {
	GetAll() ([]*domains.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) GetAll() ([]*domains.Menu, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var menus []*domains.Menu
	if err := r.db.Preload("TrainingSets").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
