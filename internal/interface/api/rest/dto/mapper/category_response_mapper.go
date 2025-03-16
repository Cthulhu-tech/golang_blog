package controller_response_mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	controller_response "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/response"
)

func ToCategoryResponse(category *common.CategoryResult) *controller_response.Category {
	return &controller_response.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func ToCategoryListResponse(categories []*common.CategoryResult) *controller_response.CategoryListResponse {
	var categoryList []*controller_response.Category

	for _, category := range categories {
		categoryList = append(categoryList, ToCategoryResponse(category))
	}

	return &controller_response.CategoryListResponse{Categories: categoryList}
}
