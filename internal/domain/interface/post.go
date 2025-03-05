package interfaces

import "github.com/Cthulhu-tech/golang_blog/internal/domain/entities"

type PostRepository interface {
	Create(post *entities.Post) error
	GetByID(id string) (*entities.Post, error)
	GetAll(page, pageSize int) ([]entities.Post, int64, error)
	Delete(id string) error
}
