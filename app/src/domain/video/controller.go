package video

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	VideoService VideoService
}

func ProvideVideoController(videoService VideoService) VideoController {
	return VideoController{VideoService: videoService}
}

func (c *VideoController) Add(ctx *gin.Context) {
	var videoInput VideoInput
	if err := ctx.BindJSON(&videoInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video := c.VideoService.Save(DTOToEntity(videoInput))

	ctx.JSON(http.StatusCreated, gin.H{"message": "added", "video": EntityToDTO(video)})
	//ctx.JSON(http.StatusCreated, gin.H{"message": "added", "video": "xxx"})
}

// func (c *VideoController) Get(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
// }

func (c *VideoController) FindAll(ctx *gin.Context) {
	videos := c.VideoService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"result": EntitySetToDTO(videos)})
}
