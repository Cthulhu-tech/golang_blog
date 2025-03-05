package interfaces

import "github.com/Cthulhu-tech/golang_blog/internal/domain/entities"

type TagRepository interface {
	Create(tag *entities.Tag) error
	GetByID(id string) (*entities.Tag, error)
	GetTagsByIDs(ids []string, tags *[]entities.Tag) error
	GetAll() ([]entities.Tag, error)
}
