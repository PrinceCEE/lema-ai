package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Phone     string    `json:"phone" gorm:"unique;not null"`
	Address   Address   `json:"address" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Address struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Street    string    `json:"street" gorm:"not null"`
	City      string    `json:"city" gorm:"not null"`
	State     string    `json:"state" gorm:"not null"`
	Zipcode   string    `json:"zipcode" gorm:"not null"`
	UserID    string    `json:"user_id" gorm:"index;not null"`
	Posts     []Post    `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
