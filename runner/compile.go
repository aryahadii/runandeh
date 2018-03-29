package runner

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

type CompileError interface {
	error
}

type CppCompileError struct {
	gccOut string
}

func (e CppCompileError) Error() string {
	return fmt.Sprintf("CppCompileError: %s", e.gccOut)
}

// GetExecutableFilePath compiles source code and returns path to the
// executable app
func GetExecutableFilePath(request *RunRequest) (string, error) {
	var err error
	var binPath string

	switch request.CodeLang {
	case LangCpp:
		binPath, err = getCompiledCppBinary(request)
		if err != nil {
			return "", err
		}
	}

	return binPath, nil
}

func getCompiledCppBinary(request *RunRequest) (string, error) {
	baseFilename := fmt.Sprintf("/tmp/cpp_%d_%s", request.ID, time.Now().Format("2006_01_02T15-04-05_07-00"))

	// Write source code to file
	sourceFilename := baseFilename + ".cpp"
	ioutil.WriteFile(sourceFilename, []byte(request.AppCode), os.ModePerm)

	// Compile
	compiler := exec.Command("g++", "-o", baseFilename, sourceFilename)
	gccOutput, err := compiler.CombinedOutput()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return "", CppCompileError{gccOut: string(gccOutput)}
		}
		return "", err
	}

	return baseFilename, nil
}
