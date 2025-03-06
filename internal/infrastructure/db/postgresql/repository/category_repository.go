package repository

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	postgres "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
	db_mapper "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/mapper"
	"gorm.io/gorm"
)

type GormCategoryRepository struct {
	db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &GormCategoryRepository{db: db}
}

func (repo *GormCategoryRepository) Create(category *entities.Category) error {
	dbCategory := db_mapper.ToDBCategory(category)

	if err := repo.db.Create(dbCategory).Error; err != nil {
		return err
	}

	return nil
}

func (repo *GormCategoryRepository) GetByID(id string) (*entities.Category, error) {
	var dbCategory postgres.Category

	err := repo.db.Preload("Category").First(&dbCategory, id).Error
	if err != nil {
		return nil, err
	}

	category := db_mapper.FromDBCategory(&dbCategory)

	return category, nil
}

func (repo *GormCategoryRepository) GetAll(page, pageSize int) ([]*entities.Category, int64, error) {
	var dbCategories []postgres.Category
	var total int64

	if err := repo.db.Model(&postgres.Category{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := repo.db.Limit(pageSize).Offset(page * pageSize).Find(&dbCategories).Error; err != nil {
		return nil, 0, nil
	}

	category := make([]*entities.Category, len(dbCategories))
	for i, dbCategory := range dbCategories {
		category[i] = db_mapper.FromDBCategory(&dbCategory)
	}

	return category, total, nil
}
