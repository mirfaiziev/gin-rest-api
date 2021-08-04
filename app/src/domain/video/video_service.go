package video

type VideoService struct {
	VideoRepository VideoRepository
}

func ProvideVideoService(repo VideoRepository) VideoService {
	return VideoService{VideoRepository: repo}
}

func (c *VideoService) FindAll() []VideoEntity {
	return c.VideoRepository.FindAll()
}

func (c *VideoService) Save(video VideoEntity) VideoEntity {
	c.VideoRepository.Save(video)

	return video
}
