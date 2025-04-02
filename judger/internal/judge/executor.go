package judge

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Run(executablePath, inputPath string, timeLimit int) (string, error) {
	cmd := exec.Command(executablePath)
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return "", fmt.Errorf("读取输入文件失败: %v", err)
	}
	cmd.Stdin = bytes.NewReader(input)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	timer := time.AfterFunc(time.Duration(timeLimit)*time.Millisecond, func() {
		cmd.Process.Kill()
	})
	err = cmd.Run()
	timer.Stop()
	if err != nil {
		return "", fmt.Errorf("运行失败: %v, 输出: %s", err, output.String())
	}
	return output.String(), nil
}
