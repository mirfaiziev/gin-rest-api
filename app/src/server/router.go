package server

import (
	"rest-api-app/src/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	v1 := router.Group("api/v1")
	{
		testGroup := v1.Group("test")
		{
			testController := new(controller.TestController)
			testGroup.GET("/", testController.Get)
		}

		videoGroup := v1.Group("video")
		{
			videoContoller := new(controller.VideoController)
			videoGroup.POST("/", videoContoller.Add)
		}
	}

	return router
}
