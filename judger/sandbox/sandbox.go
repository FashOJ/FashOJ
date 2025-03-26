package sandbox

import (
	"encoding/json"
	"os"
	"os/exec"
	"syscall"
	"time"

	seccomp "github.com/seccomp/libseccomp-golang"
)

// Sandbox is the sandbox of a single submissio, to be used to judge the code.
type SandboxConfig struct {
	TimeLimit   float64
	MemoryLimit int
	ExecFile    string
	InputFile   string
	OutputFile  string
	SeccompFile string
}

// JudgeResult is the result of a single judge, to be returned to the client the code judged result.
type JudgeResult struct {
	JudgeStatus       string
	Time              float64
	RunMemory         int //kb
	CodeReturnMessage string
}

type SeccompRule struct {
	DefaultAction string   `json:"default_action"`
	Architectures []string `json:"architectures"`
	Syscalls      []struct {
		Names  []string `json:"names"`
		Action string   `json:"action"`
	} `json:"syscalls"`
}

func loadSeccompFilter(filePath string) *syscall.SysProcAttr {
	fileJSON, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer fileJSON.Close()

	var rule SeccompRule
	if err := json.NewDecoder(fileJSON).Decode(&rule); err != nil {
		return nil
	}

	// 设置默认拒绝动作
	filter, err := seccomp.NewFilter(seccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	if err != nil {
		return nil
	}

	// 添加架构支持
	for _, arch := range rule.Architectures {
		if err := filter.AddArchitecture(seccomp.ArchFromString(arch)); err != nil {
			return nil
		}
	}

	// 添加系统调用规则
	for _, call := range rule.Syscalls {
		action := seccomp.ActFromString(call.Action)
		for _, name := range call.Names {
			if err := filter.AddRule(seccomp.ScmpSyscallFromName(name), action); err != nil {
				return nil
			}
		}
	}

	// 应用默认动作
	if err := filter.SetDefaultAction(seccomp.ActFromString(rule.DefaultAction)); err != nil {
		return nil
	}

	return &syscall.SysProcAttr{
		Seccomp: filter,
	}
}

func RunInSandbox(config SandboxConfig) (JudgeResult, error) {
	cmd := exec.Command(config.ExecFile)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Seccomp: loadSeccompFilter(config.SeccompFile),
	}

	err := cmd.Run()

	if err != nil {
		return JudgeResult{
			JudgeStatus:       "Runtime Error",
			Time:              0,
			RunMemory:         0,
			CodeReturnMessage: "RE",
		}, err
	}

	doneTime := make(chan error, 1)

	go func() {
		doneTime <- cmd.Wait()
	}()

	select {
	case <-time.After(time.Duration(config.TimeLimit) * time.Millisecond):
		_ = cmd.Process.Kill()
		return JudgeResult{
			JudgeStatus:       "Time Limit Exceeded",
			Time:              config.TimeLimit,
			RunMemory:         0,
			CodeReturnMessage: "TLE",
		}, nil
	case <-doneTime:
		return JudgeResult{
			JudgeStatus:       "Accepted",
			Time:              0,
			RunMemory:         0,
			CodeReturnMessage: "AC",
		}, nil
	}

}
