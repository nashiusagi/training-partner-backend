package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"training-partner/usecases"
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
