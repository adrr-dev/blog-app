// Package repository contains all service for models
package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	newRepository := &Repository{db: db}
	return newRepository
}
