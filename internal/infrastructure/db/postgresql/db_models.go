package postgres

import "time"

type Post struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CategoryID uint      `json:"category_id"`
	Tags       []Tag     `json:"tags" gorm:"many2many:post_tags;"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Category struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

type Tag struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}
