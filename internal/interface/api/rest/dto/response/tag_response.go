package controller_response

import "time"

type Tag struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type TagListResponse struct {
	Tags []*Tag
}
