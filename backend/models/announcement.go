package models

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Title string `binding:"required"`
	Abstract string
	Content string	 `binding:"required"`
	UserID uint
	User User
}