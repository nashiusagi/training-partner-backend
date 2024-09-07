package domains

type Exercise struct {
	ExerciseId   uint `gorm:"primary_key"`
	Name         string
	RegisteredId uint
	Muscles      []Muscle `gorm:"many2many:exercise_muscles_target_to_train;joinForeignKey:ExerciseId;joinReferences:MuscleId"`
}
