package main

import (
	"fmt"
	"log"

	"github.com/Cthulhu-tech/golang_blog/internal/application/services"
	postgres_migrate_data "github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql"
	"github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/repository"
	rest "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto"
	"github.com/Cthulhu-tech/golang_blog/internal/utils"
	"github.com/labstack/echo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	utils.LoadEnv()
	connectionString := utils.GetPostgresql()

	server := "localhost:8080"
	gormDB, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	gormDB.AutoMigrate(&postgres_migrate_data.Category{})
	gormDB.AutoMigrate(&postgres_migrate_data.Tag{})
	gormDB.AutoMigrate(&postgres_migrate_data.Post{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	categoryRepo := repository.NewGormCategoryRepository(gormDB)
	// tagRepo := repository.NewGormTagRepository(gormDB)
	// postRepo := repository.NewGormPostRepository(gormDB)

	categoryService := services.NewCategoryService(categoryRepo)
	// tagService := services.NewTagService(tagRepo)
	// postService := services.NewPostService(postRepo, tagRepo)

	e := echo.New()

	rest.NewCategoryController(e, categoryService)

	fmt.Println("Successfully connected to the database")

	if err := e.Start(server); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
