package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		InputData          string
		ExpectedOutputData string
	}{
		{
			"1 2",
			"3",
		},
		{
			"199 1",
			"200",
		},
		{
			fmt.Sprintf("%v %v", strings.Repeat("9", 1000), 1),
			"1" + strings.Repeat("0", 1000),
		},
		{
			"102 201",
			"303",
		},
		{
			"50 50",
			"100",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%v", i), func(t *testing.T) {
			input := bytes.NewBuffer([]byte(test.InputData))
			output := bytes.NewBuffer(make([]byte, 0, os.Getpagesize()))

			if err := Solution(input, output, t); err != nil {
				t.Fatalf("Solution(): %v", err)
			}

			if test.ExpectedOutputData != output.String() {
				t.Fatalf("Expected output:\n\"%s\"\nActual:\n\"%s\"", test.ExpectedOutputData, output)
			}
		})
	}

}
