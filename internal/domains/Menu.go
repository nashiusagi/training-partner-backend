package domains

import "time"

type Menu struct {
	MenuId       uint `gorm:"primary_key"`
	Date         time.Time
	TrainingSets []TrainingSet `gorm:"many2many:menus_training_sets;joinForeignKey:MenuId;joinReferences:TrainingSetId"`
}
