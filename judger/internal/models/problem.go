package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ProblemId uint
	TimeLimit uint
	MemoryLimit uint
	TestCases []Testcase `gorm:"foreignKey:ProblemID"`
}