package repository

import (
	"fmt"

	"github.com/adrr-dev/blog-app/internal/domain"
)

func (r Repository) FetchPosts(userID uint) ([]domain.Post, error) {
	var posts []domain.Post
	result := r.db.Preload("User").Preload("Comments").Find(&posts, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (r Repository) NewPost(userID uint, content string) error {
	post := &domain.Post{Content: content, UserID: userID}
	result := r.db.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r Repository) RandomPosts() ([]domain.Post, error) {
	var posts []domain.Post
	result := r.db.Order("RANDOM()").Limit(5).Preload("User").Preload("Comments").Find(&posts)
	if result.Error != nil {
		return nil, fmt.Errorf("could not query for random posts: %e", result.Error)
	}

	return posts, nil
}

func (r Repository) DeletePost(postID, userID uint) error {
	result := r.db.Delete(&domain.Post{}, "id = ? AND user_id = ?", postID, userID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
