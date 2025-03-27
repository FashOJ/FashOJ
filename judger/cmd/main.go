package main

import (
	"FashOJ/Judger/internal/judge"
	"fmt"
)

func main() {
	
	sourceCPP := "../testTEMP/sourceCode/testCPP.cpp"
	// sourcePython := "../testTEMP/sourceCode/testPython.cpp"
	exec := "../testTEMP/execTemp/test"
	input := "../testTEMP/testCase/input1.txt"
	answer := "../testTEMP/outputCase/output1.txt"
	timeLimit := 1000 // 1s
	memoryLimit := 256

	result := judge.Judge(sourceCPP, exec, input, answer, timeLimit, int64(memoryLimit))
	fmt.Println("评测结果:", result)
}
