package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../resources/training_partner.db"), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}

func TestHelloWorld(t *testing.T) {
	db, _ := setupDB()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello World", w.Body.String())
}

func TestExercises(t *testing.T) {
	db, _ := setupDB()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exercises", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "", w.Body.String())
}

func TestFindExercise(t *testing.T) {
	db, _ := setupDB()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/exercises/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t,
		"{\"ExerciseId\":1,\"Name\":\"レッグプレス\",\"RegisteredId\":5,\"Muscles\":[{\"MuscleId\":1,\"Name\":\"大腿四頭筋\",\"BodyPartId\":7},{\"MuscleId\":2,\"Name\":\"大殿筋\",\"BodyPartId\":7},{\"MuscleId\":3,\"Name\":\"下腿三頭筋\",\"BodyPartId\":7},{\"MuscleId\":4,\"Name\":\"ハムストリングス\",\"BodyPartId\":7}]}",
		w.Body.String(),
	)
}

func TestTrainingSets(t *testing.T) {
	db, _ := setupDB()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/training_sets", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEqual(t, "", w.Body.String())
}

func TestFindTrainingSets(t *testing.T) {
	db, _ := setupDB()

	router := setupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/training_sets/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"TrainingSetId\":1,\"ExerciseId\":1,\"Weight\":95,\"Repetition\":10}", w.Body.String())
}

func TestCreateTrainingSets(t *testing.T) {
	// arrange
	db, _ := setupDB()
	router := setupRouter(db)
	w := httptest.NewRecorder()

	data := url.Values{}
	data.Set("exercise_id", "2")
	data.Set("weight", "100")
	data.Set("repetition", "100")

	req, _ := http.NewRequest(http.MethodPost, "/training_sets/create", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// act
	router.ServeHTTP(w, req)

	// asset
	assert.Equal(t, 200, w.Code)
}
