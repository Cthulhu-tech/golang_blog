package controller_response_mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	controller_response "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/dto/response"
)

func ToCategoryResponse(category *common.CategoryResult) *controller_response.CategoryResponse {
	return &controller_response.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
}

func ToCategoryListResponse(categories []*common.CategoryResult) *controller_response.CategoryListResponse {
	var categoryList []*controller_response.CategoryResponse

	for _, category := range categories {
		categoryList = append(categoryList, ToCategoryResponse(category))
	}

	return &controller_response.CategoryListResponse{Categories: categoryList}
}
