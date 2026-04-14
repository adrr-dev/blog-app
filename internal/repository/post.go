package repository

import (
	"github.com/adrr-dev/blog-app/internal/domain"
)

func (r Repository) FetchPosts(userID uint) ([]domain.Post, error) {
	var posts []domain.Post
	result := r.db.Find(&posts, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (r Repository) NewPost(userID uint, content string) error {
	user, err := r.FetchUserByID(userID)
	if err != nil {
		return err
	}
	post := &domain.Post{Content: content, UserID: userID, User: user}
	result := r.db.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
