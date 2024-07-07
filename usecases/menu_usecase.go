package usecases

import (
	"training-partner/domains"
	"training-partner/repositories"
)

type MenuUsecase interface {
	GetAll() ([]*domains.Menu, error)
	FindById(id int) (*domains.Menu, error)
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
