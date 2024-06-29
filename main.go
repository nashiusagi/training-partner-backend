package main

import (
	"training-partner/controllers"
	"training-partner/repositories"
	"training-partner/usecases"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func main() {
	db, err := gorm.Open("sqlite3", "resources/post.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Post{})

	postRepository := repositories.NewPostRepository(db)
	postUseCase := usecases.NewPostUsecase(postRepository)
	postController := controllers.NewPostController(postUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/posts", postController.GetAll)

	r.Run()
}
