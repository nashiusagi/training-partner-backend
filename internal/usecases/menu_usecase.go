package usecases

import (
	"time"
	"training-partner/internal/domains"
	"training-partner/internal/repositories"
)

type MenuUsecase interface {
	GetAll() ([]*domains.Menu, error)
	FindById(id int) (*domains.Menu, error)
	Create(date time.Time) error
}

type menuUseCase struct {
	menuRepository repositories.MenuRepository
}

func NewMenuUsecase(menuRepository repositories.MenuRepository) MenuUsecase {
	return &menuUseCase{menuRepository}
}

func (u *menuUseCase) GetAll() ([]*domains.Menu, error) {
	return u.menuRepository.GetAll()
}

func (u *menuUseCase) FindById(id int) (*domains.Menu, error) {
	return u.menuRepository.FindById(id)
}

func (u *menuUseCase) Create(date time.Time) error {
	return u.menuRepository.Create(date)
}
