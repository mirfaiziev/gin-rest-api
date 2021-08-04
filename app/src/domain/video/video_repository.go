package video

import (
	"gorm.io/gorm"
)

type VideoRepository struct {
	DB *gorm.DB
}

func ProvideVideoRepository(DB *gorm.DB) VideoRepository {
	return VideoRepository{DB: DB}
}

func (c *VideoRepository) FindAll() []VideoEntity {
	var videos []VideoEntity
	c.DB.Find(&videos)

	return videos
}

func (c *VideoRepository) Save(video VideoEntity) VideoEntity {
	c.DB.Save(&video)

	return video
}
