package repositories

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) interfaces.TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) Create(tag *entities.Tag) error {
	return r.db.Create(tag).Error
}

func (r *tagRepository) GetByID(id string) (*entities.Tag, error) {
	var tag entities.Tag
	err := r.db.First(&tag, id).Error
	return &tag, err
}

func (r *tagRepository) GetAll() ([]entities.Tag, error) {
	var tags []entities.Tag
	err := r.db.Find(&tags).Error
	return tags, err
}

func (r *tagRepository) GetTagsByIDs(ids []string, tags *[]entities.Tag) error {
	return r.db.Where("id IN ?", ids).Find(tags).Error
}
