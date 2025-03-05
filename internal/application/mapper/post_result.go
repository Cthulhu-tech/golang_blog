package mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

func NewPostResultFromValidatedEntity(post *entities.Post) *common.PostResult {
	if post == nil {
		return nil
	}

	var tagResults []*common.TagResult
	for _, tag := range post.Tags {
		tagResults = append(tagResults, &common.TagResult{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return &common.PostResult{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		CategoryID: post.CategoryID,
		Tags:       tagResults,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}

func NewPostResultFromEntity(post *entities.Post, totalRecords int64) *common.PostResult {
	if post == nil {
		return nil
	}

	var tagResults []*common.TagResult
	for _, tag := range post.Tags {
		tagResults = append(tagResults, &common.TagResult{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return &common.PostResult{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		CategoryID: post.CategoryID,
		Tags:       tagResults,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}
