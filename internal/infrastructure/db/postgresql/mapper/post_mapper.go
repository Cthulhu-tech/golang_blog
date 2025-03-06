package db_mapper

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
)

func ToDBPost(post *entities.Post) *postgres.Post {
	var tags []postgres.Tag

	for _, tag := range post.Tags {
		tags = append(tags, postgres.Tag{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt,
		})
	}

	var p = &postgres.Post{
		ID:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		CategoryID: post.CategoryID,
		Tags:       tags,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}

	return p
}

func FromDBPost(dbPost *postgres.Post) *entities.Post {
	var tags []entities.Tag

	for _, tag := range dbPost.Tags {
		tags = append(tags, entities.Tag{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt,
		})
	}
	var p = &entities.Post{
		ID:         dbPost.ID,
		Title:      dbPost.Title,
		Content:    dbPost.Content,
		CategoryID: dbPost.CategoryID,
		Tags:       tags,
		CreatedAt:  dbPost.CreatedAt,
		UpdatedAt:  dbPost.UpdatedAt,
	}

	return p
}
