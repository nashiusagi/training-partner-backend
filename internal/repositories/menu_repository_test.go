package repositories_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMenuRepositoryGetAll(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	menuRepository := repositories.NewMenuRepository(mockDB)

	t.Run("正常に値を取得できた場合、Menuを全て返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus`")).
			WillReturnRows(sqlmock.
				NewRows([]string{"menu_id", "date"}).
				AddRow(uint(1), time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus_training_sets` WHERE `menus_training_sets`.`menu_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"menu_id", "training_set_id", "count"}).
				AddRow(uint(1), uint(1), uint(3)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
				AddRow(uint(1), uint(1), uint(95), uint(10)))

		menus, err := menuRepository.GetAll()

		assert.Equal(t, nil, err)
		assert.Equal(t, time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC), menus[0].Date)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Menus: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus`")).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := menuRepository.GetAll()

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Menus: %v", err)
		}
	})
}

func TestMenuRepositoryFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	menuRepository := repositories.NewMenuRepository(mockDB)

	t.Run("正常に値を取得できた場合、対応するMenuを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus` WHERE `menus`.`menu_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"menu_id", "date"}).
				AddRow(uint(1), time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus_training_sets` WHERE `menus_training_sets`.`menu_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"menu_id", "training_set_id", "count"}).
				AddRow(uint(1), uint(1), uint(3)))
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
				AddRow(uint(1), uint(1), uint(95), uint(10)))

		menu, err := menuRepository.FindById(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC), menu.Date)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Menus: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `menus` WHERE `menus`.`menu_id` = ?")).
			WithArgs(1).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := menuRepository.FindById(1)

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Menus: %v", err)
		}
	})
}

func TestMenuRepositoryCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	menuRepository := repositories.NewMenuRepository(mockDB)

	t.Run("正常にMenuを作成できた場合、nilを返す", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `menus` (`date`) VALUES (?)")).WithArgs(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = menuRepository.Create(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

		assert.Equal(t, nil, err)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Create Menu: %v", err)
		}
	})

	t.Run("正常にMenuを作成できなかった場合、作成がされずエラーを返す", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `menus` (`date`) VALUES (?)")).WithArgs(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)).WillReturnError(fmt.Errorf("some errors"))
		mock.ExpectRollback()

		err = menuRepository.Create(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Create Menu: %v", err)
		}
	})
}
