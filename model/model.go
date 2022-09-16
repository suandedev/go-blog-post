package model

import "gorm.io/gorm"

type M map[string]interface{}

type Page struct {
	Title string
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
}

type User struct {
	gorm.Model
	Name     string
	Username string
	Password string
	Post     []Post
}
