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

type postService struct {
	postRepo interfaces.PostRepository
	tagRepo  interfaces.TagRepository
}

func NewPostService(postRepo interfaces.PostRepository, tagRepo interfaces.TagRepository) app_interfaces.PostService {
	return &postService{
		postRepo: postRepo,
		tagRepo:  tagRepo,
	}
}

func (s *postService) CreatePost(title, content string, categoryID uint, tagIDs []string) (*entities.Post, error) {
	tags, err := s.tagRepo.GetTagsByIDs(tagIDs)
	if err != nil {
		return nil, err
	}

	post := entities.NewPost(title, content, categoryID, tags)

	if err := s.postRepo.Create(post); err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) GetPostByID(id string) (*query.PostyQueryResult, error) {
	utils.LogInfo(fmt.Sprintf("Get post with ID: %s", id))

	post, err := s.postRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	queryResult := &query.PostyQueryResult{
		Result: mapper.NewPostResultFromValidatedEntity(post),
	}

	return queryResult, nil
}

func (s *postService) GetAllPosts(page, pageSize int) ([]query.PostyQueryListResult, error) {
	utils.LogInfo(fmt.Sprintf("Get all post with page: %d and pageSize: %d", page, pageSize))

	postList, totalRecords, err := s.postRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, err
	}

	var queryListResult []query.PostyQueryListResult

	for _, post := range postList {
		queryListResult = append(queryListResult, query.PostyQueryListResult{
			Result: []*common.PostResult{mapper.NewPostResultFromEntity(post, totalRecords)},
		})
	}

	return queryListResult, nil
}

func (s *postService) DeletePost(id string) error {
	utils.LogInfo(fmt.Sprintf("Delete post with id: %s", id))
	return s.postRepo.Delete(id)
}
