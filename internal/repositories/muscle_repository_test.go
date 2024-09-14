package repositories_test

import (
	"regexp"
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMuscleRepositoryGetAll(t *testing.T) {
	// Arrange
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	muscleRows := sqlmock.
		NewRows([]string{"muscle_id", "name", "body_part_id"}).
		AddRow(uint(1), "大腿四頭筋", uint(7))
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `muscles`")).
		WillReturnRows(muscleRows)

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
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	muscleRows := sqlmock.
		NewRows([]string{"muscle_id", "name", "body_part_id"}).
		AddRow(uint(1), "大腿四頭筋", uint(7))
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
		WithArgs(1).
		WillReturnRows(muscleRows)

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
