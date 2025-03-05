package entities

import (
	"fmt"
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/google/uuid"
)

type Post struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID uint      `json:"category_id"`
	Tags       []Tag     `json:"tags" gorm:"many2many:post_tags;"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewPost(title, content string, categoryID uint, tags []Tag) *Post {
	var postID = uuid.New().String()

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
