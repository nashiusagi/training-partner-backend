package usecases_test

import (
	"testing"
	"training-partner/internal/domains"
	"training-partner/internal/usecases"
	"training-partner/internal/usecases/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTrainigSetUseCaseGetAll(t *testing.T) {
	mockTrainingSetRepository := new(mocks.TrainingSetRepository)

	t.Run("success", func(t *testing.T) {
		mockTrainingSets := []*domains.TrainingSet{
			{
				TrainingSetId: uint(1),
				ExerciseId:    uint(1),
				Weight:        uint(95),
				Repetition:    uint(10),
			},
		}

		mockTrainingSetRepository.On("GetAll", mock.Anything).Return(mockTrainingSets, nil).Once()

		trainingSetUseCase := usecases.NewTrainingSetUsecase(mockTrainingSetRepository)

		trainingSets, err := trainingSetUseCase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, trainingSets)
		assert.Equal(t, trainingSets[0].TrainingSetId, uint(1))
	})
}

func TestTrainigSetUseCaseFindById(t *testing.T) {
	mockTrainingSetRepository := new(mocks.TrainingSetRepository)

	t.Run("success", func(t *testing.T) {
		mockTrainingSet := domains.TrainingSet{
			TrainingSetId: uint(1),
			ExerciseId:    uint(1),
			Weight:        uint(95),
			Repetition:    uint(10),
		}

		mockTrainingSetRepository.On("FindById", 1).Return(&mockTrainingSet, nil).Once()

		trainingSetUseCase := usecases.NewTrainingSetUsecase(mockTrainingSetRepository)

		trainingSet, err := trainingSetUseCase.FindById(1)

		assert.NoError(t, err)
		assert.Equal(t, trainingSet.TrainingSetId, uint(1))
	})
}
