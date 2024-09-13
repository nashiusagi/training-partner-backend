package repositories_test

import (
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
