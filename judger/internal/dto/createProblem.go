package dto

type CreateProblem struct {
	ProblemId uint `json:"problem_id"`
	TimeLimit uint `json:"time_limit"`
	MemoryLimit uint `json:"memory_limit"`
}