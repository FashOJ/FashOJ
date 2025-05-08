package sandbox

import (
	"fmt"
	"syscall"
	// "golang.org/x/sys/unix"
)

func setupSeccomp() error {
	// 没办法啊没办法啊，草泥马能不能正常点
	return nil
}

func setupResourceLimits(config *SandboxConfig) error {
	rlimit := syscall.Rlimit{
		Cur: uint64(config.TimeLimit) * 1000000, // us
		Max: uint64(config.TimeLimit) * 1000000,
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_CPU, &rlimit); err != nil {
		return fmt.Errorf("设置 CPU 时间限制失败: %v", err)
	}

	rlimit = syscall.Rlimit{
		Cur: uint64(config.OutputLimit) * 1024, // 字节
		Max: uint64(config.OutputLimit) * 1024,
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_FSIZE, &rlimit); err != nil {
		return fmt.Errorf("设置输出大小限制失败: %v", err)
	}

	rlimit = syscall.Rlimit{
		Cur: uint64(config.ProcessLimit),
		Max: uint64(config.ProcessLimit),
	}
	if err := syscall.Setrlimit(7, &rlimit); err != nil { // 7 const -》 RLIMIT_NPROC 
		return fmt.Errorf("设置进程数量限制失败: %v", err)
	}

	return nil
}
