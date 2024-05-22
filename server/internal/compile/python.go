package compile

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// TLEError is a custom error type for time limit exceeded
type TLEError struct{}

func (e *TLEError) Error() string {
	return "execution timeout: time limit exceeded"
}

func ExecutePythonCode(pythonCode string) (string, error) {
	tmpfile, err := ioutil.TempFile("", "script-*.py")
	if err != nil {
		return err.Error(), fmt.Errorf("could not create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(pythonCode)); err != nil {
		return err.Error(), fmt.Errorf("could not write to temporary file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		return err.Error(), fmt.Errorf("could not close temporary file: %v", err)
	}

	// Create a context with a 1-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Execute the Python script using the Python interpreter with the context
	cmd := exec.CommandContext(ctx, "python3", tmpfile.Name())
	output, err := cmd.CombinedOutput()

	// Check if the context deadline was exceeded
	if ctx.Err() == context.DeadlineExceeded {
		return "execution timeout: time limit exceeded", &TLEError{}
	}

	// Return the output and any error encounteredy
	if err != nil {
		return err.Error(), fmt.Errorf("execution error: %v", err)
	}

	return string(output), nil
}
