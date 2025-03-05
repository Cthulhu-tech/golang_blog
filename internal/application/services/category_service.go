package services

import (
	"fmt"

	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	app_interfaces "github.com/Cthulhu-tech/golang_blog/internal/application/interface"
	"github.com/Cthulhu-tech/golang_blog/internal/application/mapper"
	"github.com/Cthulhu-tech/golang_blog/internal/application/query"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	"github.com/Cthulhu-tech/golang_blog/internal/utils"
)

type categoryService struct {
	categoryRepo interfaces.CategoryRepository
}

func NewCategoryService(categoryRepo interfaces.CategoryRepository) app_interfaces.CategoryService {
	return &categoryService{categoryRepo}
}

func (s *categoryService) CreateCategory(name string) (*entities.Category, error) {
	category := entities.NewCategory(name)

	err := s.categoryRepo.Create(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) GetCategoryByID(id string) (*query.CategoryQueryResult, error) {
	utils.LogInfo(fmt.Sprintf("Get category with id: %s", id))

	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	queryResult := &query.CategoryQueryResult{
		Result: mapper.NewProductResultFromValidatedEntity(category),
	}

	return queryResult, nil
}

func (s *categoryService) GetAllCategories(page, pageSize int) ([]query.CategoryQueryListResult, error) {
	utils.LogInfo(fmt.Sprintf("Get all categories with page: %d and pageSize: %d", page, pageSize))
	categoryList, totalRecords, err := s.categoryRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	var queryListResult []query.CategoryQueryListResult

	for _, category := range categoryList {
		queryListResult = append(queryListResult, query.CategoryQueryListResult{
			Result: []*common.CategoryResult{mapper.NewCategoryResultFromEntity(&category, totalRecords)},
		})
	}

	return queryListResult, nil
}
