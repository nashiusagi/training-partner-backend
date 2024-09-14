package repositories_test

import (
	"fmt"
	"regexp"
	"testing"
	"training-partner/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTrainingSetRepositoryGetAll(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)

	t.Run("正常に値を取得できた場合、TrainingSetを全て返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets`")).
			WillReturnRows(sqlmock.
				NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
				AddRow(uint(1), uint(1), uint(95), uint(10)).
				AddRow(uint(2), uint(1), uint(85), uint(10)))

		training_sets, err := trainingSetRepository.GetAll()

		assert.Equal(t, nil, err)
		assert.Equal(t, uint(95), training_sets[0].Weight)
		assert.Equal(t, uint(85), training_sets[1].Weight)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find Training Sets: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets`")).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := trainingSetRepository.GetAll()

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find TrainingSets: %v", err)
		}
	})
}

func TestTrainingSetRepositoryFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)

	t.Run("正常に値を取得できた場合、対応するTrainingSetを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
			WithArgs(1).
			WillReturnRows(sqlmock.
				NewRows([]string{"training_set_id", "exercise_id", "weight", "repetition"}).
				AddRow(uint(1), uint(1), uint(95), uint(10)))

		trainingSet, err := trainingSetRepository.FindById(1)

		assert.Equal(t, nil, err)
		assert.Equal(t, uint(1), trainingSet.TrainingSetId)
		assert.Equal(t, uint(95), trainingSet.Weight)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find TrainingSets: %v", err)
		}
	})

	t.Run("正常に値を取得できない場合、エラーを返す", func(t *testing.T) {
		mock.
			ExpectQuery(
				regexp.QuoteMeta("SELECT * FROM `training_sets` WHERE `training_sets`.`training_set_id` = ?")).
			WithArgs(1).
			WillReturnError(fmt.Errorf("some errors"))

		_, err := trainingSetRepository.FindById(1)

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Find TrainingSets: %v", err)
		}
	})
}

func TestTrainingSetRepositoryCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	trainingSetRepository := repositories.NewTrainingSetRepository(mockDB)

	t.Run("正常にTrainingSetを作成できた場合、nilを返す", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `training_sets` (`exercise_id`,`weight`,`repetition`) VALUES (?,?,?)")).WithArgs(uint(1), uint(105), uint(10)).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = trainingSetRepository.Create(uint(1), uint(105), uint(10))

		assert.Equal(t, nil, err)
		if err != nil {
			t.Fatal(err)
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Create TrainingSet: %v", err)
		}
	})

	t.Run("正常にTrainingSetを作成できなかった場合、作成がされずエラーを返す", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO `training_sets` (`exercise_id`,`weight`,`repetition`) VALUES (?,?,?)")).WithArgs(uint(1), uint(105), uint(10)).WillReturnError(fmt.Errorf("some errors"))
		mock.ExpectRollback()

		err = trainingSetRepository.Create(uint(1), uint(105), uint(10))

		if err == nil {
			t.Errorf("error is expected, but error not occured")
		}
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Test Create Training Set: %v", err)
		}
	})
}
