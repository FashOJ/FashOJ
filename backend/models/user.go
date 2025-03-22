package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string
	Right    int       `gorm:"default:0"`                           // 0 表示普通用户，1表示管理员
	Problems []Problem `json:"problems" gorm:"foreignKey:AuthorID"` // 一对多关系
}
