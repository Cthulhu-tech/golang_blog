package controller_response

import "time"

type Post struct {
	ID         string
	Title      string
	Content    string
	CategoryID uint
	Tags       []*Tag
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PostListResponse struct {
	Posts []*Post
}
