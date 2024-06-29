package controllers

import (
	"net/http"
	"strconv"
	"training-partner/usecases"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postUsecase usecases.PostUsecase
}

func NewPostController(postUsecase usecases.PostUsecase) *PostController {
	return &PostController{postUsecase}
}

func (c *PostController) GetAll(ctx *gin.Context) {
	posts, err := c.postUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	post, err := c.postUsecase.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, post)
}
