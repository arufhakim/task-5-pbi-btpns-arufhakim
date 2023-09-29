package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title    string `form:"title"`
	UserID   uint   `form:"user_id"`
	Caption  string `form:"caption"`
	PhotoURL string `form:"photo_url"`
	User     User   `gorm:"foreignKey:UserID"`
}
