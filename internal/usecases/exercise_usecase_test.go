package usecases_test

import (
	"testing"
	"training-partner/internal/domains"
	"training-partner/internal/usecases"
	"training-partner/internal/usecases/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestExerciseUseCaseGetAll(t *testing.T) {
	mockExerciseRepository := new(mocks.ExerciseRepository)

	t.Run("正常に値を取得できる", func(t *testing.T) {
		mockExercises := []*domains.Exercise{
			{
				ExerciseId:   uint(1),
				Name:         "スクワット",
				RegisteredId: uint(100),
				Muscles: []domains.Muscle{
					{
						MuscleId:   uint(2),
						Name:       "筋肉",
						BodyPartId: uint(1000),
					},
				},
			},
		}

		mockExerciseRepository.On("GetAll", mock.Anything).Return(mockExercises, nil).Once()

		exerciseUseCase := usecases.ExerciseUsecase(mockExerciseRepository)

		exercises, err := exerciseUseCase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, exercises)
		assert.Equal(t, exercises[0].Name, "スクワット")
	})
}

func TestExerciseUseCaseFindById(t *testing.T) {
	mockExerciseRepository := new(mocks.ExerciseRepository)

	t.Run("正常に値を取得できる", func(t *testing.T) {
		mockExercise := domains.Exercise{
			ExerciseId:   uint(1),
			Name:         "レッグプレス",
			RegisteredId: uint(100),
			Muscles: []domains.Muscle{
				{
					MuscleId:   uint(2),
					Name:       "筋肉",
					BodyPartId: uint(1000),
				},
			},
		}

		mockExerciseRepository.On("FindById", 1).Return(&mockExercise, nil).Once()

		exerciseUseCase := usecases.ExerciseUsecase(mockExerciseRepository)

		exercise, err := exerciseUseCase.FindById(1)

		assert.NoError(t, err)
		assert.NotNil(t, exercise)
		assert.Equal(t, exercise.Name, "レッグプレス")
	})
}
