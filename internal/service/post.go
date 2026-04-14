package service

import (
	"fmt"

	"github.com/adrr-dev/blog-app/internal/domain"
)

func (s Service) FetchPosts(userID uint) ([]domain.Post, error) {
	posts, err := s.repo.FetchPosts(userID)
	if err != nil {
		return nil, err
	}
	return posts, err
}

func (s Service) NewPost(userID uint, content string) error {
	err := s.repo.NewPost(userID, content)
	if err != nil {
		return fmt.Errorf("could not create post: %e", err)
	}
	return nil
}
