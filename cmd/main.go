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

	menuRepository := repositories.NewMenuRepository(db)
	menuUseCase := usecases.NewMenuUsecase(menuRepository)
	menuController := controllers.NewMenuController(menuUseCase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.GET("/menus", menuController.GetAll)
	r.GET("/menus/:id", menuController.FindById)

	if err := r.Run(); err != nil {
		return
	}
}
