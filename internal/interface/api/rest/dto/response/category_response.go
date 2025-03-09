package controler_response

import "time"

type CategoryResponse struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

type CategoryListResponse struct {
	Categories []*CategoryResponse
}
