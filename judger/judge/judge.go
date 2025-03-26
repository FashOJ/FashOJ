package judge

import (
	"FashOJ/Judger/sandbox"
	"FashOJ/Judger/compiler"
)

func Judge(codeFilePath string, language string, comfig sandbox.SandboxConfig) (sandbox.JudgeResult, error) {
	var err error
	var codeOutput string

	switch language {
	case "python3":
		codeOutput, err = compiler.CompliePython3(codeFilePath)
		if err != nil {
			return sandbox.JudgeResult{
				JudgeStatus:       "Compile Error",
				Time:              0,
				RunMemory:         0,
				CodeReturnMessage: "CE\n" + codeOutput,
			}, err
		}	
	}
	codeResult, err := sandbox.RunInSandbox(comfig)
	if err != nil {
		return codeResult, err
	}

	return codeResult, nil
}