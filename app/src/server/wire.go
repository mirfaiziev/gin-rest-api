//+build wireinject

package server

import (
	"rest-api-app/src/database"
	"rest-api-app/src/domain/video"

	"github.com/google/wire"
)

func initVideoController() video.VideoController {
	wire.Build(database.GetDB, video.ProvideVideoRepository, video.ProvideVideoService, video.ProvideVideoController)

	return video.VideoController{}
}
