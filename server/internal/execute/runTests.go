package execute

import (
	"fmt"
	"lms/db/tasks"
)

func RunTests(tests []tasks.Test, pythonCode string, n int) (int, int, error) {
	passed := 0
	total := len(tests)

	for i := 0; i < n; i++ {
		for _, test := range tests {
			// Execute Python code for each test input
			output, err := ExecutePythonCode(pythonCode, test.Input)
			if err != nil {
				return 0, 0, fmt.Errorf("error executing Python code: %v", err)
			}

			if output == test.Output {
				passed++
			}
		}
	}

	return passed, total * n, nil
}
