package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	TestCases []TestCase
}