package entities

import (
	"fmt"
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/application/command"
	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/google/uuid"
)

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
}

func NewCategory(categoryCommand *command.CreateCategoryCommand) *Category {
	var categoryId = uuid.New().String()

	utils.LogInfo(fmt.Sprintf("Creating category with ID: %s", categoryId))

	return &Category{
		ID:        categoryId,
		Name:      categoryCommand.Name,
		CreatedAt: time.Now(),
	}
}
