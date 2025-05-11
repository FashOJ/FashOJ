package main

import (
	"FashOJ/Judger/internal/judge"
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(filename))
	sourceCPP := filepath.Join(baseDir, "testTEMP", "sourceCode", "testCPP.cpp")
	exec := filepath.Join(baseDir, "testTEMP", "execTemp", "test")
	input := filepath.Join(baseDir, "testTEMP", "testCase", "input1.txt")
	answer := filepath.Join(baseDir, "testTEMP", "outputCase", "output1.txt")
	timeLimit := 1000 // 1s

	result := judge.Judge(sourceCPP, exec, input, answer, timeLimit)
	fmt.Println("评测结果:", result)
}