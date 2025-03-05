package postgres

import (
	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", utils.GetPostgresql())
}
