package db_mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
)

func ToDBTag(tag *entities.Tag) *postgres.Tag {
	var c = &postgres.Tag{
		ID:        tag.ID,
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
	}

	return c
}

func FromDBTag(dbTag *postgres.Tag) *entities.Tag {
	var c = &entities.Tag{
		ID:        dbTag.ID,
		Name:      dbTag.Name,
		CreatedAt: dbTag.CreatedAt,
	}

	return c
}
