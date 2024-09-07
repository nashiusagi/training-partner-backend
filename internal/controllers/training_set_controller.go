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

func (c *TrainingSetController) Create(ctx *gin.Context) {
	stringExerciseId := ctx.PostForm("exercise_id")
	exerciseId, _ := strconv.ParseUint(stringExerciseId, 10, 64)
	stringWeight := ctx.PostForm("weight")
	weight, _ := strconv.ParseUint(stringWeight, 10, 64)
	stringRepetition := ctx.PostForm("repetition")
	repetition, _ := strconv.ParseUint(stringRepetition, 10, 64)

	if exerciseId == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "exercise Id must greater than 0"})
		return
	}
	if weight == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "weight must greater than 0"})
		return
	}
	if repetition == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "repetition must greater than 0"})
		return
	}

	err := c.trainingSetUsecase.Create(uint(exerciseId), uint(weight), uint(repetition))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
