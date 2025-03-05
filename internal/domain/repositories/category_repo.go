package repositories

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) interfaces.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *entities.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id string) (*entities.Category, error) {
	var category entities.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) GetAll(page, pageSize int) ([]entities.Category, int64, error) {
	var category []entities.Category
	var totalRecords int64

	if err := r.db.Model(&entities.Category{}).Count(&totalRecords).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&category).Error; err != nil {
		return nil, 0, err
	}

	return category, totalRecords, nil
}
