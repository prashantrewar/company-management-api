package db

import (
	"company-management-api/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := os.Getenv("DSN")
    if dsn == "" {
        log.Fatal("DSN environment variable not set")
    }

    fmt.Println("Connecting to database with DSN:", dsn)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Error getting database: %v\n", err)
    }

    if err = sqlDB.Ping(); err != nil {
        log.Fatalf("Error connecting to database: %v\n", err)
    }

    fmt.Println("Connected to database")

    // Automatically migrate your models (creates tables if they do not exist)
    err = DB.AutoMigrate(&models.User{}, &models.Customer{}, &models.Billing{}, &models.Payroll{})
    if err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}
