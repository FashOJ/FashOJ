package judge

import (
	"fmt"
	"os"
	"strings"
)

func Compare(output, answerPath string) (bool, error) {
	expected, err := os.ReadFile(answerPath)
	if err != nil {
		return false, fmt.Errorf("读取outputCase失败: %v", err)
	}
	if strings.TrimSpace(output) == strings.TrimSpace(string(expected)) {
		return true, nil
	}
	return false, nil
}
