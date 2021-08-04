package server

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	v1 := router.Group("api/v1")
	{
		videoGroup := v1.Group("video")
		{
			videoContoller := initVideoController()
			videoGroup.POST("/", videoContoller.Add)
			videoGroup.GET("/", videoContoller.FindAll)
		}
	}

	return router
}
