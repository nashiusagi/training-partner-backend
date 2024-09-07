package main

import (
	"training-partner/internal/controllers"
	"training-partner/internal/repositories"
	"training-partner/internal/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := gorm.Open(sqlite.Open("resources/training_partner.db"), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}

	exerciseRepository := repositories.NewExerciseRepository(db)
	exerciseUseCase := usecases.NewExerciseUsecase(exerciseRepository)
	exerciseController := controllers.NewExerciseController(exerciseUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/exercises", exerciseController.GetAll)
	r.GET("/exercises/:id", exerciseController.FindById)

	if err := r.Run(); err != nil {
		return
	}
}
