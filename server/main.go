package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func executePythonCode(pythonCode string) (string, error) {
	// Create a temporary file
	tmpfile, err := ioutil.TempFile("", "script-*.py")
	if err != nil {
		return "", fmt.Errorf("could not create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	// Write the Python code to the temporary file
	if _, err := tmpfile.Write([]byte(pythonCode)); err != nil {
		return "", fmt.Errorf("could not write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		return "", fmt.Errorf("could not close temporary file: %v", err)
	}

	// Execute the Python script using the Python interpreter
	cmd := exec.Command("python3", tmpfile.Name())
	output, err := cmd.CombinedOutput()
	// Return the output and any error encountered
	if err != nil {
		return string(output), fmt.Errorf("execution error: %v", err)
	}
	return string(output), nil
}

func executeCppCode(cppCode string) (string, error) {
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
	defer os.Remove(binFile.Name()) // clean up

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

func main() {
	pythonCode := `
print("Hello, World!")
x = 5
y = 10
print("Sum:", x + y)
`

	cppCode := `
#include <iostream>

int main() {
    std::cout << "Hello, World!" << std::endl;
    int x = 5;
    int y = 10;
    std::cout << "Sum: " << x + y << std::endl;
    return 0;
}
`

	pythonOutput, pythonErr := executePythonCode(pythonCode)
	if pythonErr != nil {
		fmt.Printf("Python Error: %v\n", pythonErr)
	}
	fmt.Printf("Python Output: %s\n", pythonOutput)

	cppOutput, cppErr := executeCppCode(cppCode)
	if cppErr != nil {
		fmt.Printf("C++ Error: %v\n", cppErr)
	}
	fmt.Printf("C++ Output: %s\n", cppOutput)
}
