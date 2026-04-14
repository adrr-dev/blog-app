// Package service contains the service logic
package service

import (
	"github.com/adrr-dev/blog-app/internal/domain"
	"gorm.io/gorm"
)

type Repo interface {
	CreateUser(username, password string) error
	FetchUser(username, password string) (*domain.User, error)
	FetchUserByID(id uint) (*domain.User, error)

	NewPost(userID uint, content string) error
	FetchPosts(userID uint) ([]domain.Post, error)
}

type Service struct {
	db   *gorm.DB
	repo Repo
}

func NewService(db *gorm.DB, repo Repo) *Service {
	newService := &Service{db: db, repo: repo}
	return newService
}
