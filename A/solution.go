package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	if err := Solution(os.Stdin, os.Stdout); err != nil {
		 panic(err)
	}
}

func Solution(r io.Reader, w io.Writer) error {
	var result int
	var value int

	for {
		_, err := fmt.Fscanln(r, &value)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		result ^= value
	}

	_, err := fmt.Fprintln(w, result)
	return err
}
