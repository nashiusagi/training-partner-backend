package controllers

import (
	"net/http"
	"strconv"
	"training-partner/internal/usecases"

	"github.com/gin-gonic/gin"
)

type ExerciseController struct {
	exerciseUsecase usecases.ExerciseUsecase
}

func NewExerciseController(exerciseUsecase usecases.ExerciseUsecase) *ExerciseController {
	return &ExerciseController{exerciseUsecase}
}

func (c *ExerciseController) GetAll(ctx *gin.Context) {
	exercises, err := c.exerciseUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercises)
}

func (c *ExerciseController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	exercise, err := c.exerciseUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exercise)
}
