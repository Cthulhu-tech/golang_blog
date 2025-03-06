package interfaces

import "github.com/Cthulhu-tech/golang_blog/internal/domain/entities"

type CategoryRepository interface {
	Create(category *entities.Category) error
	GetByID(id string) (*entities.Category, error)
	GetAll(page, pageSize int) ([]*entities.Category, int64, error)
}
