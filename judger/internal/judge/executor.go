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
			cmd.Process.Kill()
			cmd.Process.Release()
		}
	})
	defer timer.Stop()
	done := make(chan error, 1)
	go func() { done <- cmd.Run() }()
	select {
	case err = <-done: 
	case <-time.After(time.Duration(timeLimit+100) * time.Millisecond): 
		return "", fmt.Errorf("执行超时")
	}

	// err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Run Failed: %v, Output: %s", err, output.String())
	}
	return output.String(), nil
}
