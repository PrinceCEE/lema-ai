package models

type Post struct {
	ID        string `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id" gorm:"index;not null"`
	Title     string `json:"title" gorm:"type:text;not null"`
	Body      string `json:"body" gorm:"type:text;not null"`
	CreatedAt string `json:"created_at"`
}
