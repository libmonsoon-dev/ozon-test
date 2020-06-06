package main

import (
	"testing"
)

func TestSolution(t *testing.T) {

	result, err := Solution()
	if err != nil {
		t.Fatalf("Solution(): %v", err)
	}

	expected := Result{12, "testName"}
	if len(result) != 1 || result[0] != expected {
		for _, row := range result {
			t.Logf("%+v\n", row)
		}
		t.Fatalf("Expected: %+v\n", expected)

	}

}