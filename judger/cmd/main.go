package main

import (
	"FashOJ/Judger/internal/judge"
	"fmt"
)

func main() {
	
	source := "../testTEMP/sourceCode/test.cpp"
	exec := "../testTEMP/execTemp/test"
	input := "../testTEMP/testCase/input1.txt"
	answer := "../testTEMP/outputCase/output1.txt"
	timeLimit := 1000 // 1s

	result := judge.Judge(source, exec, input, answer, timeLimit)
	fmt.Println("评测结果:", result)
}
