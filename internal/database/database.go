// Package database initiats the database
package database

import (
	"github.com/adrr-dev/blog-app/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB(dataFile string) (*gorm.DB, error) {
	DB, err := gorm.Open(sqlite.Open(dataFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = DB.AutoMigrate(&domain.User{}, &domain.Post{}, &domain.Comment{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
