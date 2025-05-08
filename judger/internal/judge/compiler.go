package judge

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CompileCpp17(sourcePath, execPath string) error {
	execDir := filepath.Dir(execPath)
	if err := os.MkdirAll(execDir, 0755); err != nil {
		return fmt.Errorf("mkdir failed: %v", err)
	}

	cmd := exec.Command("g++", "-std=c++17", sourcePath, "-o", execPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("编译失败: %v\n%s", err, string(output))
	}
	return nil
}
