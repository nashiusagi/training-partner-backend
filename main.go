package main

import (
	"training-partner/controllers"
	"training-partner/repositories"
	"training-partner/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string
	Body  string
}

func main() {
	db, err := gorm.Open(sqlite.Open("resources/post.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&Post{}); err != nil {
		return
	}

	postRepository := repositories.NewPostRepository(db)
	postUseCase := usecases.NewPostUsecase(postRepository)
	postController := controllers.NewPostController(postUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/posts", postController.GetAll)
	r.GET("/posts/:id", postController.FindById)

	if err := r.Run(); err != nil {
		return
	}
}
