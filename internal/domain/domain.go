// Package domain contains the models
package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Posts    []Post
	Comments []Comment
}

type Post struct {
	gorm.Model
	Content  string
	UserID   uint
	User     *User
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	User    *User
	PostID  uint
	Post    *Post
}
