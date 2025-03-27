package judge

import (
	"fmt"
	"os"
)

func Judge(sourcePath, execPath, inputPath, answerPath string, timeLimit int, memoryLimit int64) string {
	if err := CompileCpp17(sourcePath, execPath); err != nil {
		return fmt.Sprintf("编译错误: %v", err)
	}
	fmt.Println("编译好了")
	output, err := Run(execPath, inputPath, timeLimit, memoryLimit)
	if err != nil {
		return fmt.Sprintf("运行错误: %v", err)
	}
	fmt.Println("运行好了")
	match, err := Compare(output, answerPath)
	if err != nil {
		return fmt.Sprintf("答案错误: %v", err)
	}

	fmt.Println("评测好了")

	defer os.Remove(execPath)

	if match {
		return "ACCEPTED"
	} else {
		return "WRONG ANSWER"
	}
}
