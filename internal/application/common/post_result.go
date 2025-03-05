package common

import "time"

type PostResult struct {
	ID         string
	Title      string
	Content    string
	CategoryID uint
	Tags       []*TagResult
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
