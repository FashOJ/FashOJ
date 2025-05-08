package sandbox

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func setMemoryLimit(pid int, limitKB int) error {
	cgroupPath := "/sys/fs/cgroup/memory"
	if _, err := os.Stat(cgroupPath); os.IsNotExist(err) {
		return fmt.Errorf("cgroup memory 子系统不可用??")
	}
	sandboxCgroup := filepath.Join(cgroupPath, fmt.Sprintf("FashOJ_sandbox_%d", pid))
	if err := os.Mkdir(sandboxCgroup, 0755); err != nil {
		return fmt.Errorf("make cgroup failed: %v", err)
	}
	limitBytes := limitKB * 1024
	if err := os.WriteFile(
		filepath.Join(sandboxCgroup, "memory.limit_in_bytes"),
		[]byte(strconv.Itoa(limitBytes)),
		0644,
	); err != nil {
		return fmt.Errorf("set mem limit failed: %v", err)
	}
	if err := os.WriteFile(
		filepath.Join(sandboxCgroup, "memory.swappiness"),
		[]byte("0"),
		0644,
	); err != nil {
		return fmt.Errorf("don't use swap failed: %v", err)
	}
	if err := os.WriteFile(
		filepath.Join(sandboxCgroup, "tasks"),
		[]byte(strconv.Itoa(pid)),
		0644,
	); err != nil {
		return fmt.Errorf("add process to cgroup failed: %v", err)
	}

	return nil
}

func getMemoryUsage(pid int) (int, error) {
	cgroupPath := "/sys/fs/cgroup/memory"
	// 修改为与 setMemoryLimit 中相同的大小写
	sandboxCgroup := filepath.Join(cgroupPath, fmt.Sprintf("FashOJ_sandbox_%d", pid))

	maxUsageBytes, err := os.ReadFile(filepath.Join(sandboxCgroup, "memory.max_usage_in_bytes"))
	if err != nil {
		return 0, fmt.Errorf("read mem use status failed: %v", err)
	}

	// KB
	maxUsage, err := strconv.Atoi(string(maxUsageBytes))
	if err != nil {
		return 0, fmt.Errorf("parse mem use status failed: %v", err)
	}

	return maxUsage / 1024, nil
}

// 清理
func cleanupCgroup(pid int) error {
	cgroupPath := "/sys/fs/cgroup/memory"
	sandboxCgroup := filepath.Join(cgroupPath, fmt.Sprintf("FashOJ_sandbox_%d", pid))

	if err := os.WriteFile(
		filepath.Join(cgroupPath, "tasks"),
		[]byte(strconv.Itoa(pid)),
		0644,
	); err != nil {
		return fmt.Errorf("move out process in cgroup failed: %v", err)
	}

	if err := os.Remove(sandboxCgroup); err != nil {
		return fmt.Errorf("delete cgroup failed: %v", err)
	}

	return nil
}
