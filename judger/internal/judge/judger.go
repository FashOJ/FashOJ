package judge

import "fmt"

func Judge(sourcePath, execPath, inputPath, answerPath string, timeLimit int) string {
	if err := Compile(sourcePath, execPath); err != nil {
		return fmt.Sprintf("编译错误: %v", err)
	}
	output, err := Run(execPath, inputPath, timeLimit)
	if err != nil {
		return fmt.Sprintf("运行错误: %v", err)
	}
	match, err := Compare(output, answerPath)
	if err != nil {
		return fmt.Sprintf("答案错误: %v", err)
	}

	if match {
		return "ACCEPTED"
	} else {
		return "WRONG ANSWER"
	}
}
