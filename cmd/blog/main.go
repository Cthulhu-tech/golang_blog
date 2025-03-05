package main

import (
	"fmt"
	"log"

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

	fmt.Println("Successfully connected to the database")

}
