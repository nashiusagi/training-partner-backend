package domains

type Menu struct {
	MenuId       uint `gorm:"primary_key"`
	Name         string
	RegisteredId uint
	Muscles      []Muscle `gorm:"many2many:menu_muscles_target_to_train;joinForeignKey:MenuId;joinReferences:MenuId"`
}
