package judge

import (
	"fmt"
	"os"
	"strings"
)

func Compare(output, answerPath string) (bool, error) {
	// 移除调试输出
	// fmt.Println(output)
	
	expected, err := os.ReadFile(answerPath)
	if err != nil {
		return false, fmt.Errorf("读取outputCase失败: %v", err)
	}
	
	// 规范化输出和答案（去除空白字符和换行符）
	normalizedOutput := strings.TrimSpace(output)
	normalizedExpected := strings.TrimSpace(string(expected))
	
	// 比较规范化后的字符串
	if normalizedOutput == normalizedExpected {
		return true, nil
	}
	
	return false, nil
}
