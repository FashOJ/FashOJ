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
		MemoryLimit:  262144,    
		TimeLimit:    1000,      
		OutputLimit:  64 * 1024,  // 暂时的，后面改成可修改变量
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
	
	if s.WorkDir == "" {
		s.WorkDir = filepath.Dir(s.ExecPath)
	}
	cmd.Dir = s.WorkDir

	if s.InputPath != "" {
		input, err := os.Open(s.InputPath)
		if err != nil {
			return nil, fmt.Errorf("Open FUCK input file failed: %v", err)
		}
		defer input.Close()
		cmd.Stdin = input
	}
	
	cmd.Stdout = tempOutput
	cmd.Stderr = tempError

	cmd.SysProcAttr = &syscall.SysProcAttr{}
	
	// Damn it Namespaces 
	setupNamespaces(cmd.SysProcAttr)
	
	startTime := time.Now()
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("cnm启动程序失败: %v", err)
	}
	
	pid := cmd.Process.Pid
	
	if err := setMemoryLimit(pid, s.MemoryLimit); err != nil {
		cmd.Process.Kill()
		return nil, fmt.Errorf("setting failed: %v", err)
	}
	
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()
	
	var timedOut bool
	select {
	case err = <-done:
		// 程序正常结束
		/**
 *                             _ooOoo_
 *                            o8888888o
 *                            88" . "88
 *                            (| -_- |)
 *                            O\  =  /O
 *                         ____/`---'\____
 *                       .'  \\|     |//  `.
 *                      /  \\|||  :  |||//  \
 *                     /  _||||| -:- |||||-  \
 *                     |   | \\\  -  /// |   |
 *                     | \_|  ''\---/''  |   |
 *                     \  .-\__  `-`  ___/-. /
 *                   ___`. .'  /--.--\  `. . __
 *                ."" '<  `.___\_<|>_/___.'  >'"".
 *               | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *               \  \ `-.   \_ __\ /__ _/   .-` /  /
 *          ======`-.____`-.___\_____/___.-`____.-'======
 *                             `=---='
 *          ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
 *                     佛祖保佑        永无BUG
*/
	case <-time.After(time.Duration(s.TimeLimit) * time.Millisecond):
		// Fuck it, 超时了妈的
		// 杀进程
		cmd.Process.Kill()
		timedOut = true
		err = <-done
	}
	
	executionTime := int(time.Since(startTime).Milliseconds())
	
	memoryUsage, memErr := getMemoryUsage(pid)
	if memErr != nil {
		memoryUsage = 0
	}
	
	cleanupCgroup(pid)
	
	tempOutput.Seek(0, 0)
	outputBytes, err := os.ReadFile(tempOutput.Name())
	if err != nil {
		return nil, fmt.Errorf("read output failed: %v", err)
	}
	
	tempError.Seek(0, 0)
	errorBytes, err := os.ReadFile(tempError.Name())
	if err != nil {
		return nil, fmt.Errorf("read error output failed: %v", err)
	}

	
	result := &SandboxResult{
		Time:        executionTime,
		Memory:      memoryUsage,
		Output:      string(outputBytes),
		ErrorOutput: string(errorBytes),
	}

	if timedOut {
		result.Status = "Time Limit Exceeded"
		return result, nil
	}
	
	if memoryUsage > s.MemoryLimit {
		result.Status = "Memory Limit Exceeded"
		return result, nil
	}

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

	return result, nil
}