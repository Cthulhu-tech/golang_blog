package db_mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
)

func ToDBCategory(category *entities.Category) *postgres.Category {
	var c = &postgres.Category{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	return c
}

func FromDBCategory(dbCategory *postgres.Category) *entities.Category {
	var c = &entities.Category{
		ID:        dbCategory.ID,
		Name:      dbCategory.Name,
		CreatedAt: dbCategory.CreatedAt,
	}

	return c
}
