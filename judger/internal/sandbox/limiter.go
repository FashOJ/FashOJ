package sandbox

import (
	"fmt"
	"syscall"
)


// ms, mb
func SetResourceLimits(timeLimit int, memoryLimit int64) error {
	cpuLimit := syscall.Rlimit{
		Cur: uint64(timeLimit / 1000),
		Max: uint64(timeLimit / 1000),
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_CPU, &cpuLimit); err != nil {
		return fmt.Errorf("error: %v", err)
	}
	memLimit := syscall.Rlimit{
		Cur: uint64(memoryLimit * 1024 * 1024),
		Max: uint64(memoryLimit * 1024 * 1024),
	}
	if err := syscall.Setrlimit(syscall.RLIMIT_AS, &memLimit); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}
