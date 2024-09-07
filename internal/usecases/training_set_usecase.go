package usecases

import (
	"training-partner/internal/domains"
	"training-partner/internal/repositories"
)

type TrainingSetUsecase interface {
	GetAll() ([]*domains.TrainingSet, error)
	FindById(id int) (*domains.TrainingSet, error)
}

type trainingSetUseCase struct {
	trainingSetRepository repositories.TrainingSetRepository
}

func NewTrainingSetUsecase(trainingSetRepository repositories.TrainingSetRepository) TrainingSetUsecase {
	return &trainingSetUseCase{trainingSetRepository}
}

func (u *trainingSetUseCase) GetAll() ([]*domains.TrainingSet, error) {
	return u.trainingSetRepository.GetAll()
}

func (u *trainingSetUseCase) FindById(id int) (*domains.TrainingSet, error) {
	return u.trainingSetRepository.FindById(id)
}
