package judge

import (
	"fmt"
	"os"
	"os/exec"
)

// Compile 负责编译 C++ 代码
func CompileCpp17(sourcePath, outputPath string) error {
	cmd := exec.Command("g++", sourcePath, "-o", outputPath, "-O2", "-std=c++17", "-Wall", "-Wextra")
	cmd.Stderr = os.Stderr // 捕获编译错误

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("编译失败: %v", err)
	}
	return nil
}

