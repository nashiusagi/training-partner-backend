package domains

type Menu struct {
	MenuId       uint `gorm:"primary_key"`
	Name         string
	RegisteredId uint
}
