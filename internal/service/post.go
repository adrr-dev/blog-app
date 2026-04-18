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

func (s Service) RandomPosts() ([]domain.Post, error) {
	posts, err := s.repo.RandomPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s Service) DeletePost(postID, userID uint) error {
	err := s.repo.DeletePost(postID, userID)
	if err != nil {
		return err
	}
	return nil
}
