package models

import "gorm.io/gorm"

type Testcase struct {
	gorm.Model
	ProblemID uint    // Foreign key for Problem
	InputFileID uint  // Foreign key for Input File
	OutputFileID uint // Foreign key for Output File
	
	InputFile File `gorm:"foreignKey:InputFileID"`
	OutputFile File `gorm:"foreignKey:OutputFileID"`
}