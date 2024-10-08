package usecases_test

import (
	"testing"
	"time"
	"training-partner/internal/domains"
	"training-partner/internal/usecases"
	"training-partner/internal/usecases/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMenuUseCaseGetAll(t *testing.T) {
	mockMenuRepository := new(mocks.MenuRepository)
	menuUseCase := usecases.NewMenuUsecase(mockMenuRepository)

	t.Run("正常に値を取得できる", func(t *testing.T) {
		mockMenus := []*domains.Menu{
			{
				MenuId: uint(1),
				Date:   time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
				TrainingSets: []domains.TrainingSet{
					{
						TrainingSetId: uint(1),
						ExerciseId:    uint(1),
						Weight:        95,
						Repetition:    10,
					},
				},
			},
		}
		mockMenuRepository.On("GetAll", mock.Anything).Return(mockMenus, nil).Once()

		menus, err := menuUseCase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, menus)
		assert.Equal(t, uint(1), menus[0].MenuId)
	})
}

func TestMenuUseCaseFindById(t *testing.T) {
	mockMenuRepository := new(mocks.MenuRepository)
	menuUseCase := usecases.NewMenuUsecase(mockMenuRepository)

	t.Run("正常に値を取得できる", func(t *testing.T) {
		mockMenu := domains.Menu{
			MenuId: uint(1),
			Date:   time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			TrainingSets: []domains.TrainingSet{
				{
					TrainingSetId: uint(1),
					ExerciseId:    uint(1),
					Weight:        95,
					Repetition:    10,
				},
			},
		}
		mockMenuRepository.On("FindById", 1).Return(&mockMenu, nil).Once()

		menu, err := menuUseCase.FindById(1)

		assert.NoError(t, err)
		assert.NotNil(t, menu)
		assert.Equal(t, uint(1), menu.MenuId)
	})
}

func TestMenuUseCaseCreate(t *testing.T) {
	mockMenuRepository := new(mocks.MenuRepository)
	menuUseCase := usecases.NewMenuUsecase(mockMenuRepository)

	t.Run("正常にMenuを作成できる", func(t *testing.T) {
		var dummyDate = time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC)
		mockMenuRepository.On("Create", dummyDate).Return(nil).Once()

		err := menuUseCase.Create(dummyDate)

		assert.NoError(t, err)
	})
}
