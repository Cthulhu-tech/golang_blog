package app_interfaces

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/query"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

type TagService interface {
	CreateTag(name string) (*entities.Tag, error)
	GetTagByID(id string) (*query.TagQueryResult, error)
	GetAllTags() ([]query.TagQueryListResult, error)
}
