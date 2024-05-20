package compile

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExecuteCppCode(cppCode string) (string, error) {
	// Create a temporary file for the C++ source code
	srcFile, err := ioutil.TempFile("", "program-*.cpp")
	if err != nil {
		return "", fmt.Errorf("could not create temporary source file: %v", err)
	}
	defer os.Remove(srcFile.Name()) // clean up

	// Write the C++ code to the temporary file
	if _, err := srcFile.Write([]byte(cppCode)); err != nil {
		return "", fmt.Errorf("could not write to temporary source file: %v", err)
	}
	if err := srcFile.Close(); err != nil {
		return "", fmt.Errorf("could not close temporary source file: %v", err)
	}

	// Create a temporary file for the compiled binary
	binFile, err := ioutil.TempFile("", "program-*.out")
	if err != nil {
		return "", fmt.Errorf("could not create temporary binary file: %v", err)
	}
	defer func() {
		binFile.Close()
		os.Remove(binFile.Name())
	}()

	// Compile the C++ code
	compileCmd := exec.Command("g++", srcFile.Name(), "-o", binFile.Name())
	compileOutput, err := compileCmd.CombinedOutput()
	if err != nil {
		return string(compileOutput), fmt.Errorf("compilation error: %v", err)
	}

	// Execute the compiled binary
	execCmd := exec.Command(binFile.Name())
	execOutput, err := execCmd.CombinedOutput()
	if err != nil {
		return string(execOutput), fmt.Errorf("execution error: %v", err)
	}

	// Return the output from the execution
	return string(execOutput), nil
}
