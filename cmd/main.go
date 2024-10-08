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

func setupRouter(db *gorm.DB) *gin.Engine {
	exerciseRepository := repositories.NewExerciseRepository(db)
	exerciseUseCase := usecases.NewExerciseUsecase(exerciseRepository)
	exerciseController := controllers.NewExerciseController(exerciseUseCase)
	trainingSetRepository := repositories.NewTrainingSetRepository(db)
	trainingSetUsecase := usecases.NewTrainingSetUsecase(trainingSetRepository)
	trainingSetController := controllers.NewTrainingSetController(trainingSetUsecase)
	menuRepository := repositories.NewMenuRepository(db)
	menuUsecase := usecases.NewMenuUsecase(menuRepository)
	menuController := controllers.NewMenuController(menuUsecase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/exercises", exerciseController.GetAll)
	r.GET("/exercises/:id", exerciseController.FindById)
	r.GET("/training_sets", trainingSetController.GetAll)
	r.GET("/training_sets/:id", trainingSetController.FindById)
	r.POST("/training_sets/create", trainingSetController.Create)
	r.GET("/menus", menuController.GetAll)
	r.GET("/menus/:id", menuController.FindById)
	r.POST("/menus/create", menuController.Create)

	return r
}

func main() {
	db, err := gorm.Open(sqlite.Open("resources/training_partner.db"), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}
	r := setupRouter(db)

	if err := r.Run(); err != nil {
		return
	}
}
