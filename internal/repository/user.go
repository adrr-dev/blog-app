package repository

import (
	"fmt"

	"github.com/adrr-dev/blog-app/internal/domain"
)

func (r Repository) CreateUser(username, password string) error {
	newUser := &domain.User{Username: username, Password: password}

	result := r.db.Create(newUser)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r Repository) FetchUser(username, password string) (*domain.User, error) {
	var user domain.User
	result := r.db.Preload("Posts").Preload("Comments").First(&user, "username = ? AND password = ?", username, password)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found, username or password incorrect: %e", result.Error)
	}

	return &user, nil
}

func (r Repository) FetchUserByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.db.Preload("Posts").Preload("Comments").First(&user, "ID = ?", id)
	if result.Error != nil {
		return nil, fmt.Errorf("user not found, username or password incorrect: %e", result.Error)
	}

	return &user, nil
}
