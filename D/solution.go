package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func main() {
	if err := Solution(os.Stdin, os.Stdout, nil); err != nil {
		panic(err)
	}
}

func Solution(r io.Reader, w io.Writer, logger Logger) error {
	var (
		a, b string
	)

	if _, err := fmt.Fscanln(r, &a, &b); err != nil {
		return fmt.Errorf("fmt.Fscanln(r, &a, &b): %w", err)
	}

	result := make([]string, 0, 1000)

	var (
		i, aCurrentDigit, bCurrentDigit, nextPosition int
		err                                           error
		aHasDigits, bHasDigits                        bool
	)

	for {
		aIndex := len(a) - i - 1
		bIndex := len(b) - i - 1

		if aIndex >= 0 {
			aHasDigits = true
			aCurrentDigit, err = strconv.Atoi(string(a[aIndex]))
			if err != nil {
				return fmt.Errorf("int parsing error (i: %v, a[aIndex]: %v): %w", i, string(a[aIndex]), err)
			}
		} else {
			aHasDigits = false
			aCurrentDigit = 0
		}

		if bIndex >= 0 {
			bHasDigits = true
			bCurrentDigit, err = strconv.Atoi(string(b[bIndex]))
			if err != nil {
				return fmt.Errorf("int parsing error (i: %v, b[bIndex]: %v): %w", i, string(b[bIndex]), err)
			}
		} else {
			bHasDigits = false
			bCurrentDigit = 0
		}

		if logger != nil {
			logger.Logf(
				"i: %v, aCurrentDigit: %v, bCurrentDigit: %v, nextPosition: %v, aHasDigits: %b, bHasDigits: %b",
				i, aCurrentDigit, bCurrentDigit, nextPosition, aHasDigits, bHasDigits,
			)
		}

		if !aHasDigits && !bHasDigits && nextPosition == 0 {
			break
		}

		currentSum := aCurrentDigit + bCurrentDigit + nextPosition

		nextPosition = currentSum / 10
		result = append(result, strconv.Itoa(currentSum%10))

		i++
	}

	if logger != nil {
		logger.Logf("result: %v", result)
	}

	// Reverse slice
	resultLen := len(result)
	for i := 0; i < resultLen/2; i++ {
		result[i], result[resultLen-i-1] = result[resultLen-i-1], result[i]
	}

	if logger != nil {
		logger.Logf("result: %v", result)
	}

	output := strings.Join(result, "")
	if len(output) == 0 {
		output = "0"
	}

	if _, err := fmt.Fprint(w, output); err != nil {
		return fmt.Errorf("fmt.Fprint(w, %v): %w", result, err)
	}

	return nil
}
