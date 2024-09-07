package repositories

import (
	"training-partner/internal/domains"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:generate mockery --name TrainingSetRepository
type TrainingSetRepository interface {
	GetAll() ([]*domains.TrainingSet, error)
	FindById(id int) (*domains.TrainingSet, error)
	Create(exercise_id uint, weight uint, repetition uint) error
}

type trainingSetRepository struct {
	db *gorm.DB
}

func NewTrainingSetRepository(db *gorm.DB) TrainingSetRepository {
	return &trainingSetRepository{db}
}

func (r *trainingSetRepository) GetAll() ([]*domains.TrainingSet, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var trainingSets []*domains.TrainingSet
	if err := r.db.Find(&trainingSets).Error; err != nil {
		return nil, err
	}
	return trainingSets, nil
}

func (r *trainingSetRepository) FindById(id int) (*domains.TrainingSet, error) {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var trainingSet *domains.TrainingSet
	if err := r.db.Find(&trainingSet, id).Error; err != nil {
		return nil, err
	}
	return trainingSet, nil
}

func (r *trainingSetRepository) Create(exerciseId uint, weight uint, repetition uint) error {
	r.db.Logger = r.db.Logger.LogMode(logger.Info)
	var trainingSet = domains.TrainingSet{
		ExerciseId: exerciseId,
		Weight:     weight,
		Repetition: repetition,
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&trainingSet).Error; err != nil {
			return err
		}

		return nil
	})
}
