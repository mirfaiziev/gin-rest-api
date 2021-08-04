package video

import "github.com/oklog/ulid/v2"

type VideoEntity struct {
	ID          ulid.ULID `gorm:"primarykey;unique;type:uuid;column:id"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
}
