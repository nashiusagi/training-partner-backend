package domains

type Muscle struct {
	MuscleId   uint `gorm:"primary_key"`
	Name       string
	BodyPartId uint
}
