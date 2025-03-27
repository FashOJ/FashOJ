package sandbox

import (
	"fmt"
	"syscall"

	"github.com/seccomp/libseccomp-golang"
)

func ApplySeccomp() error {
	filter, err := seccomp.NewFilter(seccomp.ActKillThread)
	if err != nil {
		return fmt.Errorf("Create seccomp Failed: %v", err)
	}
	allowedSyscalls := []uint32{
		syscall.SYS_READ,
		syscall.SYS_WRITE,
		syscall.SYS_EXIT,
		syscall.SYS_FORK,
		syscall.SYS_EXECVE,
		syscall.SYS_BRK,
		syscall.SYS_MMAP,
		syscall.SYS_MUNMAP,
		syscall.SYS_ARCH_PRCTL,
		syscall.SYS_CLOSE,
	}
	for _, sc := range allowedSyscalls {
		if err := filter.AddRule(seccomp.ScmpSyscall(sc), seccomp.ActAllow); err != nil {
			return fmt.Errorf("添加 seccomp 规则失败: %v", err)
		}
	}
	return filter.Load()
}
