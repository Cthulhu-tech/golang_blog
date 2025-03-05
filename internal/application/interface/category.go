package app_interfaces

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/query"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

type CategoryService interface {
	CreateCategory(name string) (*entities.Category, error)
	GetCategoryByID(id string) (*query.CategoryQueryResult, error)
	GetAllCategories(page, pageSize int) ([]query.CategoryQueryListResult, error)
}
