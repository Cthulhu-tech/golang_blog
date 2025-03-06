package entities

import (
	"fmt"
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/google/uuid"
)

type Post struct {
	ID         string
	Title      string
	Content    string
	CategoryID uint
	Tags       []Tag
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewPost(title, content string, categoryID uint, tags []Tag) *Post {
	postID := uuid.New().String()

	utils.LogInfo(fmt.Sprintf("Creating post with ID: %s", postID))

	return &Post{
		ID:         postID,
		Title:      title,
		Content:    content,
		CategoryID: categoryID,
		Tags:       tags,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
