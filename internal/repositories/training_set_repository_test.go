package repositories_test

import (
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTrainingSetRepositoryGetAll(t *testing.T) {
	// Arrange
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.
		NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
		AddRow(uint(1), uint(1), uint(95), uint(10)).
		AddRow(uint(1), uint(1), uint(85), uint(10))

	mock.
		ExpectQuery(
			"SELECT * FROM training_sets",
		).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)
	exercises, err := trainingSetRepository.GetAll()

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, exercises[0].Weight, uint(95))

	if err != nil {
		t.Fatal(err)
	}
}
