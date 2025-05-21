package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	InputFile File
	OutputFile File
}