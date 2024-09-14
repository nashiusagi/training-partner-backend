package repositories_test

import (
	"regexp"
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	// NOTE: sqliteのモックはできなさそうなので、mysqlを使用している
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: mockDB, SkipInitializeWithVersion: true}), &gorm.Config{})

	return db, mock, err
}

func TestExerciseRepositoryGetAll(t *testing.T) {
	// Arrange
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	exerciseRows := sqlmock.
		NewRows([]string{"exercise_id", "name", "registered_id"}).
		AddRow(uint(1), "スクワット", uint(5))
	exerciseMusclesRows := sqlmock.
		NewRows([]string{"id", "exercise_id", "muscle_id"}).
		AddRow(uint(1), uint(1), uint(1))
	muscleRows := sqlmock.
		NewRows([]string{"muscle_id", "name", "body_part_id"}).
		AddRow(uint(1), "大腿四頭筋", uint(7))
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `exercises`")).
		WillReturnRows(exerciseRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `exercise_muscles_target_to_train` WHERE `exercise_muscles_target_to_train`.`exercise_id` = ?")).
		WithArgs(1).
		WillReturnRows(exerciseMusclesRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
		WithArgs(1).
		WillReturnRows(muscleRows)

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

	exerciseRows := sqlmock.
		NewRows([]string{"exercise_id", "name", "registered_id"}).
		AddRow(uint(1), "スクワット", uint(5))
	exerciseMusclesRows := sqlmock.
		NewRows([]string{"id", "exercise_id", "muscle_id"}).
		AddRow(uint(1), uint(1), uint(1))
	muscleRows := sqlmock.
		NewRows([]string{"muscle_id", "name", "body_part_id"}).
		AddRow(uint(1), "大腿四頭筋", uint(7))

	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `exercises` WHERE `exercises`.`exercise_id` = ?")).
		WithArgs(1).
		WillReturnRows(exerciseRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `exercise_muscles_target_to_train` WHERE `exercise_muscles_target_to_train`.`exercise_id` = ?")).
		WithArgs(1).
		WillReturnRows(exerciseMusclesRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
		WithArgs(1).
		WillReturnRows(muscleRows)

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
