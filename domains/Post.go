package domains

type Post struct {
	ID    uint `gorm:"primary_key"`
	Title string
	Body  string
}
