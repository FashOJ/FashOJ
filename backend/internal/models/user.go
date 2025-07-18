package models

import "gorm.io/gorm"

// "User" class represents a user in the system.
// It contains the user's username, password, email, and permission level.
type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Avatar       string
	Password     string
	Email        string
	Permission   uint32    `gorm:"default:0"`
	Problems     []Problem `json:"problems" gorm:"foreignKey:AuthorID"` // one-to-many relationship
	Announcement []Announcement
}
