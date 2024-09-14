package controllers

import (
	"net/http"
	"strconv"
	"time"
	"training-partner/internal/usecases"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	menuUsecase usecases.MenuUsecase
}

func NewMenuController(menuUsecase usecases.MenuUsecase) *MenuController {
	return &MenuController{menuUsecase}
}

func (c *MenuController) GetAll(ctx *gin.Context) {
	menus, err := c.menuUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

func (c *MenuController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	menu, err := c.menuUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

func (c *MenuController) Create(ctx *gin.Context) {
	stringDate := ctx.PostForm("date")
	date, err := time.Parse("20060102", stringDate)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "time parse error"})
		return
	}

	err = c.menuUsecase.Create(date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
