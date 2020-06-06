package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputData = `1
2
2
1
2
3
2`
	expectedOutputData = `3
`
)

func TestSolution(t *testing.T) {
	input := bytes.NewBuffer([]byte(inputData))
	expectedOutput := bytes.NewBuffer([]byte(expectedOutputData))

	output := bytes.NewBuffer(make([]byte, 0, os.Getpagesize()))
	if err := Solution(input, output); err != nil {
		t.Fatalf("Solution() error = %v", err)
	}

	if expectedOutput.String() != output.String() {
		t.Fatalf("Expected output:\n\"%s\"\nActual:\n\"%s\"", expectedOutput, output)
	}

}
