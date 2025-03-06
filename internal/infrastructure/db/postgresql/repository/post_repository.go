package repository

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
	db_mapper "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/mapper"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) interfaces.PostRepository {
	return &GormPostRepository{db: db}
}

func (repo *GormPostRepository) Create(post *entities.Post) error {
	dbPost := db_mapper.ToDBPost(post)
	if err := repo.db.Create(dbPost).Error; err != nil {
		return err
	}

	return nil
}

func (repo *GormPostRepository) GetByID(id string) (*entities.Post, error) {
	var dbPost postgres.Post

	err := repo.db.Preload("Post").First(&dbPost, id).Error
	if err != nil {
		return nil, err
	}

	return db_mapper.FromDBPost(&dbPost), nil
}

func (repo *GormPostRepository) GetAll(page, pageSize int) ([]*entities.Post, int64, error) {
	var dbPosts []postgres.Post
	var total int64

	if err := repo.db.Model(&postgres.Post{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := repo.db.Limit(pageSize).Offset(page * pageSize).Find(&dbPosts).Error; err != nil {
		return nil, 0, nil
	}

	post := make([]*entities.Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		post[i] = db_mapper.FromDBPost(&dbPost)
	}

	return post, total, nil
}

func (repo *GormPostRepository) Delete(id string) error {
	if err := repo.db.Delete(&postgres.Post{}, id).Error; err != nil {
		return err
	}

	return nil
}
