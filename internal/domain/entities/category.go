package entities

import (
	"fmt"
	"time"

	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/google/uuid"
)

type Category struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

func NewCategory(name string) *Category {
	var categoryId = uuid.New().String()

	utils.LogInfo(fmt.Sprintf("Creating category with ID: %s", categoryId))

	return &Category{
		ID:        categoryId,
		Name:      name,
		CreatedAt: time.Now(),
	}
}
