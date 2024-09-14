package repositories_test

import (
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
	menuRows := sqlmock.
		NewRows([]string{"menu_id", "date"}).
		AddRow(uint(1), time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC))
	menuTrainingTargetRows := sqlmock.
		NewRows([]string{"menu_id", "training_set_id", "count"}).
		AddRow(uint(1), uint(1), uint(3))
	trainingSetRows := sqlmock.
		NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
		AddRow(uint(1), uint(1), uint(95), uint(10))
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `menus`")).
		WillReturnRows(menuRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `menus_training_sets` WHERE `menus_training_sets`.`menu_id` = ?")).
		WithArgs(1).
		WillReturnRows(menuTrainingTargetRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
		WithArgs(1).
		WillReturnRows(trainingSetRows)

	menuRepository := repositories.NewMenuRepository(mockDB)
	menus, err := menuRepository.GetAll()

	// TODO: mockを直す
	assert.Equal(t, nil, err)
	assert.Equal(t, time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC), menus[0].Date)
}

func TestMenuRepositoryFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	menuRows := sqlmock.
		NewRows([]string{"menu_id", "date"}).
		AddRow(uint(1), time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC))
	menuTrainingTargetRows := sqlmock.
		NewRows([]string{"menu_id", "training_set_id", "count"}).
		AddRow(uint(1), uint(1), uint(3))
	trainingSetRows := sqlmock.
		NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
		AddRow(uint(1), uint(1), uint(95), uint(10))
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `menus` WHERE `menus`.`menu_id` = ?")).
		WithArgs(1).
		WillReturnRows(menuRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `menus_training_sets` WHERE `menus_training_sets`.`menu_id` = ?")).
		WithArgs(1).
		WillReturnRows(menuTrainingTargetRows)
	mock.
		ExpectQuery(
			regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
		WithArgs(1).
		WillReturnRows(trainingSetRows)

	menuRepository := repositories.NewMenuRepository(mockDB)
	menu, err := menuRepository.FindById(1)

	assert.Equal(t, nil, err)
	assert.Equal(t, time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC), menu.Date)
}

func TestMenuRepositoryCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `menus` (`date`) VALUES (?)")).WithArgs(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC)).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	menuRepository := repositories.NewMenuRepository(mockDB)
	err = menuRepository.Create(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

	// Assert
	assert.Equal(t, nil, err)
}
