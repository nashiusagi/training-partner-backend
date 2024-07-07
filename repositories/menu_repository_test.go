package repositories_test

import (
	"regexp"
	"testing"
	"training-partner/repositories"

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

	mockDB, err := gorm.Open(sqlite.Open("../resources/training_partner.db"), &gorm.Config{})

	return mockDB, mock, err
}

func TestGetAll(t *testing.T) {
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
			"SELECT * FROM menus",
		).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	menuRepository := repositories.NewMenuRepository(mockDB)
	menus, err := menuRepository.GetAll()

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, menus[0].MenuId, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Menus: %v", err)
	// }
}

func TestFindById(t *testing.T) {
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
			`SELECT * FROM "menus" WHERE id = ?`,
		)).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	menuRepository := repositories.NewMenuRepository(mockDB)
	menu, err := menuRepository.FindById(1)

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, menu.MenuId, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Menus: %v", err)
	// }
}
