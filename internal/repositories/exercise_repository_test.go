package repositories_test

import (
	"fmt"
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
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	exerciseRepository := repositories.NewExerciseRepository(mockDB)

	t.Run("正常に値を取得できた場合、Exerciseを全て返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercises`")).
			WillReturnRows(sqlmock.
				NewRows([]string{"exercise_id", "name", "registered_id"}).
				AddRow(uint(1), "スクワット", uint(5)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercise_muscles_target_to_train` WHERE `exercise_muscles_target_to_train`.`exercise_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "exercise_id", "muscle_id"}).
				AddRow(uint(1), uint(1), uint(1)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"muscle_id", "name", "body_part_id"}).
				AddRow(uint(1), "大腿四頭筋", uint(7)))

		exercises, err := exerciseRepository.GetAll()

		assert.Equal(t, nil, err)
		assert.Equal(t, uint(1), exercises[0].ExerciseId)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Exercises: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercises`")).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := exerciseRepository.GetAll()

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Exercises: %v", err)
		}
	})
}

func TestExerciseRepositoryFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	exerciseRepository := repositories.NewExerciseRepository(mockDB)

	t.Run("正常に値を取得できた場合、対応するExerciseを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercises` WHERE `exercises`.`exercise_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"exercise_id", "name", "registered_id"}).
				AddRow(uint(1), "スクワット", uint(5)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercise_muscles_target_to_train` WHERE `exercise_muscles_target_to_train`.`exercise_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"id", "exercise_id", "muscle_id"}).
				AddRow(uint(1), uint(1), uint(1)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"muscle_id", "name", "body_part_id"}).
				AddRow(uint(1), "大腿四頭筋", uint(7)))

		exercise, err := exerciseRepository.FindById(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, "スクワット", exercise.Name)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Exercises: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `exercises` WHERE `exercises`.`exercise_id` = ?")).
			WithArgs(1).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := exerciseRepository.FindById(1)

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Exercises: %v", err)
		}
	})
}
