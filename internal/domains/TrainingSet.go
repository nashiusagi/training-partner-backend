package domains

type TrainingSet struct {
	TrainingSetId uint `gorm:"primary_key"`
	ExerciseId    uint
	Weight        uint
	Repetition    uint
}
