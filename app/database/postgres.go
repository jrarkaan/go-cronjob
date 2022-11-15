package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jrarkaan/go-cronjob/pkg/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

func Connect() (*gorm.DB, error) {
	errEnv := godotenv.Load()
	helpers.PanicIfError(errEnv)
	dbPort := os.Getenv("postgres.DB_PORT")
	dbUser := os.Getenv("postgres.DB_USER")
	dbPass := os.Getenv("postgres.DB_PASS")
	dbHost := os.Getenv("postgres.DB_HOST")
	dbName := os.Getenv("postgres.DB_NAME")
	// dsn gorm
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	fmt.Println("dsn: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helpers.PanicIfError(err)
	// pooling db
	sqlDB, errPool := db.DB()
	helpers.PanicIfError(errPool)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	return db, nil
}
