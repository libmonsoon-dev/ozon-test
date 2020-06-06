package main

import (
	"bytes"
	"os"
	"testing"
)

const (
	inputData = `5
1 7 3 4 7 9`
	expectedOutputData = `1`
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
