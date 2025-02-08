package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Username  string  `json:"username" gorm:"unique"`
	Phone     string  `json:"phone" gorm:"unique"`
	Address   Address `json:"address" gorm:"foreignKey:UserID"`
}

type Address struct {
	gorm.Model
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode string `json:"zipcode"`
	UserID  uint   `json:"user_id" gorm:"index"`
}
