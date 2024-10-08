package repositories

import (
	"time"
	"training-partner/internal/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:generate mockery --name MenuRepository
type MenuRepository interface {
	GetAll() ([]*domains.Menu, error)
	FindById(id int) (*domains.Menu, error)
	Create(date time.Time) error
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

func (r *menuRepository) FindById(id int) (*domains.Menu, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var menu *domains.Menu
	if err := r.db.Preload("TrainingSets").Find(&menu, id).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *menuRepository) Create(date time.Time) error {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var menu = domains.Menu{
		Date: date,
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&menu).Error; err != nil {
			return err
		}

		return nil
	})
}
