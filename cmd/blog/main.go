package main

import (
	"fmt"
	"log"

	"github.com/Cthulhu-tech/golang_blog/internal/application/services"
	"github.com/Cthulhu-tech/golang_blog/internal/infrastructure/db/postgresql/repository"
	rest "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto"
	controller_response "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/response"
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

	gormDB.AutoMigrate(&controller_response.Category{})

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
