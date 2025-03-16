package controller_response

import "time"

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type CategoryListResponse struct {
	Categories []*Category
}
