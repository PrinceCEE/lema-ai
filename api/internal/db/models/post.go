package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
