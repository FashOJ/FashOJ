package compiler

import (
	"os/exec"
)

// CompliePython3 is used to complie the python3 code.
func CompliePython3(codeFilePath string) (string, error) {
	cmd := exec.Command("python3", "-m", "py_complie", codeFilePath)
	codeOutput, err := cmd.CombinedOutput()
	if err != nil {
		return string(codeOutput), err
	}
	return "", nil
}