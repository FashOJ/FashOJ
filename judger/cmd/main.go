package main

import (
	"FashOJ/Judger/internal/judge"
	"fmt"
)

func main() {
	sourceCPP := "../testTEMP/sourceCode/testCPP.cpp"
	// exec := "../testTEMP/execTemp/test"
	// exec := "/home/lingbou/WorkSpace/FashOJ/judger/testTEMP/execTemp/test" // shit ！！！，记得后面把相对路径全换成绝对路径
	exec := "/home/lingbou/FashOJ/judger/testTEMP/execTemp/test" 
	///home/lingbou/WorkSpace/FashOJ/judger/testTEMP
	input := "../testTEMP/testCase/input1.txt"
	answer := "../testTEMP/outputCase/output1.txt"
	timeLimit := 200000 // 1s

	result := judge.Judge(sourceCPP, exec, input, answer, timeLimit)
	fmt.Println("评测结果:", result)
}