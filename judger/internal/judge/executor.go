package judge

import (
	"FashOJ/Judger/internal/sandbox"
	"fmt"
	// "os"
)

func Run(executablePath, inputPath string, timeLimit int) (string, error) {
	sandboxConfig := sandbox.NewSandbox()
	sandboxConfig.ExecPath = executablePath
	sandboxConfig.InputPath = inputPath
	sandboxConfig.TimeLimit = timeLimit
	
	result, err := sandboxConfig.Run()
	if err != nil {
		return "", fmt.Errorf("沙箱运行失败: %v", err)
	}
	
	switch result.Status {
	case "Accepted":
		return result.Output, nil
	case "Time Limit Exceeded":
		return "", fmt.Errorf("时间超限")
	case "Memory Limit Exceeded":
		return "", fmt.Errorf("内存超限")
	case "Runtime Error":
		return "", fmt.Errorf("运行时错误: %s", result.ErrorOutput)
	default:
		return "", fmt.Errorf("系统错误: %s", result.ErrorOutput)
	}
}
