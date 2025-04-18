package models

import "gorm.io/gorm"

// Problem Database Model
type Problem struct {
	gorm.Model
	Try        int        `gorm:"dafault:0"`
	Ac         int        `gorm:"dafault:0"`
	AuthorID   uint       `json:"author_id"`                         // 外键，指向 User 表的主键
	Author     User       `json:"author" gorm:"foreignKey:AuthorID"` // 关联 User 表
	Content    string     `json:"content"`
	Difficulty int64      `json:"difficulty"`
	ProblemID  string     `json:"problem_id" gorm:"unique"` // 唯一索引
	Title      string     `json:"title"`
	Limit      Limit      `json:"limit" gorm:"foreignKey:ProblemID"`    // 一对一关系
	Testcase   []Testcase `json:"testcases" gorm:"foreignKey:ProblemID"`
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
	Input     string `json:"input"`
	Output    string `json:"output"`
}
