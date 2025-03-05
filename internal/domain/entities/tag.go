package entities

import (
	"fmt"
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/google/uuid"
)

type Tag struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

func NewTag(name string) *Tag {
	var tagId = uuid.New().String()

	utils.LogInfo(fmt.Sprintf("Creating tag with ID: %s", tagId))

	return &Tag{
		ID:        tagId,
		Name:      name,
		CreatedAt: time.Now(),
	}
}
