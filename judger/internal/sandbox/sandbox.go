package sandbox

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

type SandboxConfig struct {
	MemoryLimit  int
	TimeLimit    int
	OutputLimit  int
	ProcessLimit int

	ExecPath   string
	InputPath  string
	OutputPath string

	WorkDir string
}

type SandboxResult struct {
	Status      string // "Accepted", "Time Limit Exceeded", "Memory Limit Exceeded", "Runtime Error", "System Error"
	ExitCode    int
	Time        int
	Memory      int
	Output      string
	ErrorOutput string
}

func NewSandbox() *SandboxConfig {
	return &SandboxConfig{
		MemoryLimit:  256,
		TimeLimit:    1000,
		OutputLimit:  64 * 1024, // 暂时的，后面改成可修改变量
		ProcessLimit: 1,
		WorkDir:      "",
	}
}

func (s *SandboxConfig) Run() (*SandboxResult, error) {
	tempOutput, err := os.CreateTemp("", "sandbox_output_")
	if err != nil {
		return nil, fmt.Errorf("make temp output file failed: %v", err)
	}
	defer os.Remove(tempOutput.Name())
	defer tempOutput.Close()

	tempError, err := os.CreateTemp("", "sandbox_error_")
	if err != nil {
		return nil, fmt.Errorf("make temp error output failed: %v", err)
	}
	defer os.Remove(tempError.Name())
	defer tempError.Close()

	cmd := exec.Command(s.ExecPath)

	cg, err := NewCgroup("test")

	if err != nil {
		return nil, fmt.Errorf("create cgroup failed: %v", err)
	}

	cg.SetMemoryLimit(s.MemoryLimit)

	if s.WorkDir == "" {
		s.WorkDir = filepath.Dir(s.ExecPath)
	}
	cmd.Dir = s.WorkDir

	// 使用hook,启用seccomp
	cmd.Env = append(cmd.Env, "LD_PRELOAD=/usr/local/lib/hook.so")

	if s.InputPath != "" {
		input, err := os.Open(s.InputPath)
		if err != nil {
			return nil, fmt.Errorf("open FUCK input file failed: %v", err)
		}
		defer input.Close()
		cmd.Stdin = input
	}

	cmd.Stdout = tempOutput
	cmd.Stderr = tempError

	cmd.SysProcAttr = &syscall.SysProcAttr{}

	// Damn it Namespaces
	setupNamespaces(cmd.SysProcAttr)

	cmd.Start()
	pid := cmd.Process.Pid
	cg.AddProc(pid)

	startTime := time.Now()

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var timedOut bool
	select {
	case err = <-done:
	case <-time.After(time.Duration(s.TimeLimit) * time.Millisecond):
		// Fuck it, 超时了妈的
		// 杀进程
		cmd.Process.Kill()
		timedOut = true
		err = <-done
	}

	result := &SandboxResult{}

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitError.ExitCode()
			result.Status = "Runtime Error"
		} else {
			result.Status = "System Error"
		}
	} else {
		result.Status = "Accepted"
		result.ExitCode = 0
	}
	result.Time = int(time.Since(startTime).Milliseconds())
	oom, err := cg.CheckOom()
	if err != nil {
		return nil, fmt.Errorf("get memory usage failed: %v", err)
	}

	cg.DelCgroup()

	tempOutput.Seek(0, 0)
	outputBytes, err := os.ReadFile(tempOutput.Name())
	result.Output = string(outputBytes)
	if err != nil {
		return nil, fmt.Errorf("read output failed: %v", err)
	}

	tempError.Seek(0, 0)
	errorBytes, err := os.ReadFile(tempError.Name())
	result.ErrorOutput = string(errorBytes)
	if err != nil {
		return nil, fmt.Errorf("read error output failed: %v", err)
	}


	if timedOut {
		result.Status = "Time Limit Exceeded"
		return result, nil
	}

	if oom {
		result.Status = "Memory Limit Exceeded"
		fmt.Println(result.Status)
		return result, nil
	}


	return result, nil
}
