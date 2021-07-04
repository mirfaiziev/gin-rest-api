package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestController struct {
}

func (c *TestController) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
}
