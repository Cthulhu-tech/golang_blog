package app_interfaces

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/command"
	"github.com/Cthulhu-tech/golang_blog/internal/application/query"
)

type CategoryService interface {
	CreateCategory(productCommand *command.CreateCategoryCommand) (*command.CreateCategoryCommandResult, error)
	GetCategoryByID(id string) (*query.CategoryQueryResult, error)
	GetAllCategories(page, pageSize int) (*query.CategoryQueryListResult, error)
}
