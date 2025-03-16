package controller_response

import "time"

type Post struct {
	ID         string
	Title      string
	Content    string
	CategoryID uint
	Tags       []*Tag `gorm:"many2many:post_tags;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PostListResponse struct {
	Posts []*Post
}
