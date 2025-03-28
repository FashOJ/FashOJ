package judge

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
	"FashOJ/Judger/internal/sandbox"
)

func Run(executablePath, inputPath string, timeLimit int, memoryLimit int64) (string, error) {
	cmd := exec.Command(executablePath)
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("read InputFile Failed: %v", err)
	}
	cmd.Stdin = bytes.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	if err := sandbox.SetResourceLimits(timeLimit, memoryLimit); err != nil {
		return "", fmt.Errorf("%v", err)
	}
	if err := sandbox.ApplySeccomp(); err != nil {
		return "", fmt.Errorf("seccomp Failed: %v", err)
	}
	timer := time.AfterFunc(time.Duration(timeLimit)*time.Millisecond, func() {
		if cmd.Process != nil {
			cmd.Process.Kill()  // 增加空指针检查
			cmd.Process.Release() // 确保释放系统资源
		}
	})
	defer timer.Stop()  // 使用defer确保定时器停止

	err = cmd.Run()
	// 移除原来的timer.Stop()
	if err != nil {
		return "", fmt.Errorf("Run Failed: %v, Output: %s", err, output.String())
	}
	return output.String(), nil
}
