package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	inputFileName  = "input.txt"
	outputFileName = "output.txt"
)

func main() {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	if err := Solution(inputFile, outputFile); err != nil {
		panic(err)
	}
}

func Solution(rOrig io.Reader, w io.Writer) error {
	var (
		target, value uint
		result        uint8
	)

	r := bufio.NewReader(rOrig)

	expectedValues := make(map[uint]struct{})

	if _, err := fmt.Fscanln(r, &target); err != nil {
		return fmt.Errorf("fmt.Fscanln(r, &target): %w", err)
	}

	for {
		_, err := fmt.Fscan(r, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("fmt.Fscan(r, &value): %w", err)
		}
		if value > target {
			continue
		}

		if _, ok := expectedValues[value]; ok {
			result = 1
			break
		} else {
			expectedValues[target-value] = struct{}{}
		}
	}

	_, err := fmt.Fprint(w, result)
	if err != nil {
		return fmt.Errorf("fmt.Fprint(w, %v): %w", result, err)
	}

	return nil
}
