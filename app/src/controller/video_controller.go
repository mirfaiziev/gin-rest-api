package controller

import (
	"net/http"
	videoDTO "rest-api-app/src/dto/video"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
}

func (c *VideoController) Add(ctx *gin.Context) {
	var videoInput videoDTO.VideoInput
	if err := ctx.BindJSON(&videoInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "added"})
}
