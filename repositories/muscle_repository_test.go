package repositories_test

import (
	"testing"
	"training-partner/repositories"

	"github.com/stretchr/testify/assert"
)

func TestMuscleRepositoryGetAll(t *testing.T) {
	// Arrange
	mockDB, _, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// Act
	muscleRepository := repositories.NewMuscleRepository(mockDB)
	muscles, err := muscleRepository.GetAll()

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, muscles[0].MuscleId, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Menus: %v", err)
	// }
}

func TestMuscleRepositoryFindById(t *testing.T) {
	// Arrange
	mockDB, _, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// Act
	muscleRepository := repositories.NewMuscleRepository(mockDB)
	muscle, err := muscleRepository.FindById(1)

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, muscle.MuscleId, uint(1))
	assert.Equal(t, muscle.Name, "大腿四頭筋")

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Menus: %v", err)
	// }
}
