package app_interfaces

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/query"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

type PostService interface {
	CreatePost(title, content string, categoryID uint, tagIDs []string) (*entities.Post, error)
	GetPostByID(id string) (*query.PostyQueryResult, error)
	GetAllPosts(page, pageSize int) ([]query.PostyQueryListResult, error)
	DeletePost(id string) error
}
