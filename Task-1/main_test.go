package main

import (
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"
)


func TestCalculateAverageGrade(t *testing.T) {
	tests := []struct {
		name     string
		grades   []float64
		expected float64
	}{
		{"All perfect scores", []float64{100, 100, 100}, 100},
		{"Mixed scores", []float64{90, 80, 70}, 80},
		{"Single score", []float64{85}, 85},
		{"Low scores", []float64{50, 60, 55}, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateAverageGrade(tt.grades)
			if result != tt.expected {
				t.Errorf("Expected average: %.2f, but got: %.2f", tt.expected, result)
			}
		})
	}
}

// Helper function to simulate user input
func simulateUserInput(input string) (func(), *bufio.Reader) {
	oldStdin := os.Stdin 

	reader := bufio.NewReader(strings.NewReader(input))
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		defer w.Close()
		w.WriteString(input)
	}()

	return func() { os.Stdin = oldStdin }, reader
}

// TestMainFunction tests the main function
func TestMainFunction(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedOutput string
	}{
		{
			name:          "Valid input",
			input:         "John Doe\n2\nMath\n85\nScience\n90\n",
			expectedOutput: "Student Name: John Doe\nSubject: Math, Grade: 85.00\nSubject: Science, Grade: 90.00\nAverage Grade: 87.50\n",
		},
		{
			name:          "Quit input",
			input:         "q\n",
			expectedOutput: "Exiting program. Goodbye!\n",
		},
		{
			name:          "Invalid number of subjects",
			input:         "Jane Doe\n-1\nq\n",
			expectedOutput: "Invalid input. Please enter a positive integer.\nExiting program. Goodbye!\n",
		},
		{
			name:          "Invalid grade",
			input:         "Mark\n2\nMath\n150\nq\n",
			expectedOutput: "Invalid input. Please enter a grade between 0 and 100.\nExiting program. Goodbye!\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output bytes.Buffer
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			done := make(chan struct{})
			go func() {
				buf := new(bytes.Buffer)
				buf.ReadFrom(r)
				output.Write(buf.Bytes())
				done <- struct{}{}
			}()

			restoreStdin, _ := simulateUserInput(tt.input)
			defer restoreStdin()

			main()

			w.Close()
			os.Stdout = oldStdout
			<-done

			if !strings.Contains(output.String(), tt.expectedOutput) {
				t.Errorf("Expected output:\n%s\nBut got:\n%s", tt.expectedOutput, output.String())
			}
		})
	}
}
