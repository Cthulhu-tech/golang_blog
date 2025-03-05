package repositories

import (
	"github.com/Cthulhu-tech/golang_blog/internal/domain/entities"
	interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) interfaces.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *entities.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) GetByID(id string) (*entities.Post, error) {
	var post entities.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll(page, pageSize int) ([]entities.Post, int64, error) {
	var posts []entities.Post
	var total int64

	err := r.db.Model(&entities.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

func (r *postRepository) Delete(id string) error {
	err := r.db.Delete(&entities.Post{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
