package sandbox

import (
	"syscall"
)

func setupNamespaces(cmd *syscall.SysProcAttr) {
	cmd.Cloneflags |= syscall.CLONE_NEWPID
	cmd.Cloneflags |= syscall.CLONE_NEWNET
	cmd.Cloneflags |= syscall.CLONE_NEWNS
	cmd.Cloneflags |= syscall.CLONE_NEWUTS
	cmd.Cloneflags |= syscall.CLONE_NEWIPC
}
