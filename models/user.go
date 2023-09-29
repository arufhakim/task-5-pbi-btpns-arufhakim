package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `form:"username" gorm:"unique;not null"`
	Email    string `form:"email" gorm:"unique;not null"`
	Password string `form:"password" gorm:"not null"`
}
