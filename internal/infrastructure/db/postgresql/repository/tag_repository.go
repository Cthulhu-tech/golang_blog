package repository

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
	db_mapper "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/mapper"
	"gorm.io/gorm"
)

type GormTagRepository struct {
	db *gorm.DB
}

func NewGormTagRepository(db *gorm.DB) interfaces.TagRepository {
	return &GormTagRepository{db: db}
}

func (repo *GormTagRepository) Create(tag *entities.Tag) error {
	dbTag := db_mapper.ToDBTag(tag)

	if err := repo.db.Create(dbTag).Error; err != nil {
		return err
	}

	return nil
}

func (repo *GormTagRepository) GetByID(id string) (*entities.Tag, error) {
	var dbTag postgres.Tag

	err := repo.db.Preload("Tag").First(&dbTag, id).Error
	if err != nil {
		return nil, err
	}

	return db_mapper.FromDBTag(&dbTag), nil
}

func (repo *GormTagRepository) GetTagsByIDs(ids []string) ([]entities.Tag, error) {
	var tags []entities.Tag

	if len(ids) == 0 {
		return tags, nil
	}

	if err := repo.db.Where("id IN ?", ids).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}

func (repo *GormTagRepository) GetAll() ([]*entities.Tag, error) {
	var dbTags []postgres.Tag
	if err := repo.db.Find(&dbTags).Error; err != nil {
		return nil, err
	}

	tags := make([]*entities.Tag, len(dbTags))
	for i, dbTag := range dbTags {
		tags[i] = db_mapper.FromDBTag(&dbTag)
	}

	return tags, nil
}
