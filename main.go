package main

import (
	"training-partner/controllers"
	"training-partner/domains"
	"training-partner/repositories"
	"training-partner/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("resources/training_partner.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&domains.Menu{}); err != nil {
		return
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
