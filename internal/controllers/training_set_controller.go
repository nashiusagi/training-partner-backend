package controllers

import (
	"net/http"
	"strconv"
	"training-partner/internal/usecases"

	"github.com/gin-gonic/gin"
)

type TrainingSetController struct {
	trainingSetUsecase usecases.TrainingSetUsecase
}

func NewTrainingSetController(trainingSetUsecase usecases.TrainingSetUsecase) *TrainingSetController {
	return &TrainingSetController{trainingSetUsecase}
}

func (c *TrainingSetController) GetAll(ctx *gin.Context) {
	trainingSets, err := c.trainingSetUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, trainingSets)
}

func (c *TrainingSetController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	trainingSet, err := c.trainingSetUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, trainingSet)
}
