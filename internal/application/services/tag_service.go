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

type tagService struct {
	tagRepo interfaces.TagRepository
}

func NewTagService(tagRepo interfaces.TagRepository) app_interfaces.TagService {
	return &tagService{tagRepo}
}

func (s *tagService) CreateTag(name string) (*entities.Tag, error) {
	tag := &entities.Tag{
		Name: name,
	}

	err := s.tagRepo.Create(tag)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *tagService) GetTagByID(id string) (*query.TagQueryResult, error) {
	utils.LogInfo(fmt.Sprintf("Get tag with id: %s", id))
	tag, err := s.tagRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	queryResult := &query.TagQueryResult{
		Result: mapper.NewTagResultFromEntity(tag),
	}

	return queryResult, nil
}

func (s *tagService) GetAllTags() ([]query.TagQueryListResult, error) {
	tagList, err := s.tagRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var queryListResult []query.TagQueryListResult

	for _, tag := range tagList {
		queryListResult = append(queryListResult, query.TagQueryListResult{
			Result: []*common.TagResult{mapper.NewTagResultFromEntity(tag)},
		})
	}

	return queryListResult, nil
}
