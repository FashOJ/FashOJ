package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Content    string    `json:"content"`
	Difficulty int64     `json:"difficulty"`
	ProblemID  string    `json:"problem_id" gorm:"unique"` // 唯一索引
	Title      string    `json:"title"`
	Examples   []Example `json:"examples" gorm:"foreignKey:ProblemID"` // 一对多关系
	Limit      Limit     `json:"limit" gorm:"foreignKey:ProblemID"`    // 一对一关系
	Testcase []Testcase `json:"testcases" gorm:"foreignKey:ProblemID"`
}

type Example struct {
	gorm.Model
	ProblemID string `json:"problem_id"` // 外键
	Input     string `json:"input"`
	Output    string `json:"output"`
}

type Limit struct {
	gorm.Model
	ProblemID   string `json:"problem_id"`   // 外键
	MemoryLimit int    `json:"memory_limit"` // KB
	TimeLimit   int    `json:"time_limit"`   // S
}

type Testcase struct {
	gorm.Model
	ProblemID string `json:"problem_id"`
	Input     string
	Output    string
}
