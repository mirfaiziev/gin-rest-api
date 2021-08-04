package video

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func DTOToEntity(videoInput VideoInput) VideoEntity {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	return VideoEntity{ID: id, Title: videoInput.Title, Description: videoInput.Description}
}

func EntityToDTO(video VideoEntity) VideoOutput {
	return VideoOutput{ID: video.ID.String(), Title: video.Title}
}

func EntitySetToDTO(videos []VideoEntity) []VideoOutput {
	result := make([]VideoOutput, len(videos))

	for i, entity := range videos {
		result[i] = EntityToDTO(entity)
	}

	return result
}
