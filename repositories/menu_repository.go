package repositories

import (
	"training-partner/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MenuRepository interface {
	GetAll() ([]*domains.Menu, error)
	FindById(id int) (*domains.Menu, error)
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
	if err := r.db.Preload("Muscles").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuRepository) FindById(id int) (*domains.Menu, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var menu *domains.Menu
	if err := r.db.Preload("Muscles").Find(&menu, id).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
