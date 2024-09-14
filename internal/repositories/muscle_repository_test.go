package repositories_test

import (
	"fmt"
	"regexp"
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMuscleRepositoryGetAll(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	muscleRepository := repositories.NewMuscleRepository(mockDB)

	t.Run("正常に値を取得できた場合、Muscleを全て返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles`")).
			WillReturnRows(sqlmock.
				NewRows([]string{"muscle_id", "name", "body_part_id"}).
				AddRow(uint(1), "大腿四頭筋", uint(7)))

		muscles, err := muscleRepository.GetAll()

		assert.Equal(t, nil, err)
		assert.Equal(t, "大腿四頭筋", muscles[0].Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Muscles: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles`")).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := muscleRepository.GetAll()

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Muscles: %v", err)
		}
	})
}

func TestMuscleRepositoryFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	muscleRepository := repositories.NewMuscleRepository(mockDB)

	t.Run("正常に値を取得できた場合、対応するMuscleを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"muscle_id", "name", "body_part_id"}).
				AddRow(uint(1), "大腿四頭筋", uint(7)))

		muscle, err := muscleRepository.FindById(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint(1), muscle.MuscleId)
		assert.Equal(t, "大腿四頭筋", muscle.Name)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Muscles: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `muscles` WHERE `muscles`.`muscle_id` = ?")).
			WithArgs(1).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := muscleRepository.FindById(1)

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Muscles: %v", err)
		}
	})
}
