package mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/application/common"
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
)

func NewTagResultFromValidatedEntity(tag *entities.Tag) *common.TagResult {
	return NewTagResultFromEntity(tag)
}

func NewTagResultFromEntity(tag *entities.Tag) *common.TagResult {
	if tag == nil {
		return nil
	}

	return &common.TagResult{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
	}
}
