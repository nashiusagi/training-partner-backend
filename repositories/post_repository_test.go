package repositories_test

import (
	"regexp"
	"testing"
	"training-partner/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	_, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	mockDB, err := gorm.Open(sqlite.Open("../resources/post.db"), &gorm.Config{})

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
			"SELECT * FROM posts",
		).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	postRepository := repositories.NewPostRepository(mockDB)
	posts, err := postRepository.GetAll()

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, posts[0].ID, uint(1))

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Posts: %v", err)
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
			`SELECT * FROM "posts" WHERE id = ?`,
		)).
		WithArgs(1).
		WillReturnRows(rows)

	// Act
	postRepository := repositories.NewPostRepository(mockDB)
	post, err := postRepository.FindById(1)

	// Assert
	assert.Equal(t, err, nil)
	assert.Equal(t, post.ID, uint(1))
	assert.Equal(t, post.Title, "title1")
	assert.Equal(t, post.Body, "body1")

	if err != nil {
		t.Fatal(err)
	}

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Find Posts: %v", err)
	// }
}
