package repositories_test

import (
	"regexp"
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	_, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	mockDB, err := gorm.Open(sqlite.Open("../../resources/training_partner.db"), &gorm.Config{})

	return mockDB, mock, err
}

func TestExerciseRepositoryGetAll(t *testing.T) {
	// Arrange
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.
		NewRows([]string{"id", "title", "body"}).
		AddRow(uint(1), "title1", "body1")

	mock.
		ExpectQuery(
			"SELECT * FROM exercises",
		).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	exerciseRepository := repositories.NewExerciseRepository(mockDB)
	exercises, err := exerciseRepository.GetAll()

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, exercises[0].ExerciseId, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Exercises: %v", err)
	// }
}

func TestExerciseRepositoryFindById(t *testing.T) {
	// Arrange
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	rows := sqlmock.
		NewRows([]string{"id", "title", "body"}).
		AddRow(uint(1), "title1", "body1")

	mock.
		ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "exercises" WHERE id = ?`,
		)).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	exerciseRepository := repositories.NewExerciseRepository(mockDB)
	exercise, err := exerciseRepository.FindById(1)

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, exercise.ExerciseId, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Exercises: %v", err)
	// }
}
