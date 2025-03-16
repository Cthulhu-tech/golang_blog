package request

import (
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/application/command"
)

type CreateCategoryRequest struct {
	Name      string    `json:"Name"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func (req *CreateCategoryRequest) ToCreateCategoryCommand() (*command.CreateCategoryCommand, error) {
	return &command.CreateCategoryCommand{
		Name:      req.Name,
		CreatedAt: req.CreatedAt,
	}, nil
}
