package repositories_test

import (
	"regexp"
	"testing"
	"time"
	"training-partner/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestMenuRepositoryGetAll(t *testing.T) {
	mockDB, _, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	menuRepository := repositories.NewMenuRepository(mockDB)
	menus, err := menuRepository.GetAll()

	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(menus))
	assert.Equal(t, time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), menus[0].Date)
}

func TestMenuRepositoryFindById(t *testing.T) {
	mockDB, _, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	menuRepository := repositories.NewMenuRepository(mockDB)
	menu, err := menuRepository.FindById(1)

	assert.Equal(t, nil, err)
	assert.Equal(t, time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), menu.Date)
}

func TestMenuRepositoryCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()

	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO menus")).WithArgs(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

	menuRepository := repositories.NewMenuRepository(mockDB)
	err = menuRepository.Create(time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC))

	// Assert
	assert.Equal(t, nil, err)
}
