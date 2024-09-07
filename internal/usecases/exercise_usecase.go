package usecases

import (
	"training-partner/internal/domains"
	"training-partner/internal/repositories"
)

type ExerciseUsecase interface {
	GetAll() ([]*domains.Exercise, error)
	FindById(id int) (*domains.Exercise, error)
}

type exerciseUseCase struct {
	exerciseRepository repositories.ExerciseRepository
}

func NewMenuUsecase(exerciseRepository repositories.ExerciseRepository) ExerciseUsecase {
	return &exerciseUseCase{exerciseRepository}
}

func (u *exerciseUseCase) GetAll() ([]*domains.Exercise, error) {
	return u.exerciseRepository.GetAll()
}

func (u *exerciseUseCase) FindById(id int) (*domains.Exercise, error) {
	return u.exerciseRepository.FindById(id)
}
