package main

import (
	"fmt"
	"log"

	"github.com/Cthulhu-tech/golang_blog/internal/application/services"
	"github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/repository"
	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := utils.GetPostgresql()

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	categoryRepo := repository.NewGormCategoryRepository(gormDB)
	tagRepo := repository.NewGormTagRepository(gormDB)
	postRepo := repository.NewGormPostRepository(gormDB)

	categoryService := services.NewCategoryService(categoryRepo)
	tagService := services.NewTagService(tagRepo)
	postService := services.NewPostService(postRepo, tagRepo)

	fmt.Println("Successfully connected to the database")

}
