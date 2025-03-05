package mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

func NewProductResultFromValidatedEntity(category *entities.Category) *common.CategoryResult {
	if category == nil {
		return nil
	}

	return &common.CategoryResult{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func NewCategoryResultFromEntity(category *entities.Category, totalRecords int64) *common.CategoryResult {

	if category == nil {
		return nil
	}

	return &common.CategoryResult{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}
