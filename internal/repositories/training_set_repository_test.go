package repositories_test

import (
	"regexp"
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
			regexp.QuoteMeta("SELECT * FROM `training_sets`")).
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

func TestTrainingSetRepositoryFindById(t *testing.T) {
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
			regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)
	trainingSet, err := trainingSetRepository.FindById(1)

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, trainingSet.TrainingSetId, uint(1))
	assert.Equal(t, trainingSet.Weight, uint(95))

	if err != nil {
		t.Fatal(err)
	}
}

func TestTrainingSetRepositoryCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `training_sets` (`exercise_id`,`weight`,`repetition`) VALUES (?,?,?)")).WithArgs(uint(1), uint(105), uint(10)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)
	err = trainingSetRepository.Create(uint(1), uint(105), uint(10))

	// Assert
	assert.Equal(t, err, nil)
}
