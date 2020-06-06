package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Logger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

func asChan(chanId int, logger Logger, slice ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range slice {
			duration := time.Duration(rand.Intn(10)) * time.Millisecond
			time.Sleep(duration)
			if logger != nil {
				logger.Logf("[chan %v] send %v\t(sleep duration %v)\n", chanId, v, duration)
			}
			c <- v
		}
		close(c)
	}()
	return c
}

func asSlice(ch <-chan int, expectedLength int) []int {
	result := make([]int, 0, expectedLength)

	for i := 0; i < expectedLength; i++ {
		val, ok := <-ch
		if !ok {
			panic("Channel closed!")
		}
		result = append(result, val)
	}
	return result
}

func TestMerge2Channels(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Logf("seed: %v\n", seed)
	rand.Seed(seed)

	tests := []struct {
		FirstInput     []int
		SecondInput    []int
		N              int
		ExpectedResult []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1,0},
			11,
			[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			16,
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			11,
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99},
			100,
			[]int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#%v", i+1), func(t *testing.T) {
			result := make(chan int)
			Merge2Channels(
				func(arg int) int {
					duration := time.Duration(rand.Intn(1500)) * time.Millisecond
					time.Sleep(duration)
					t.Logf("processed %v\t(sleep duration %v)\n", arg, duration)
					return arg
				},
				asChan(1, t, test.FirstInput...),
				asChan(2, t, test.SecondInput...),
				result,
				test.N,
			)

			resultSlice := asSlice(result, test.N)

			if expectedResultLen, resultSliceLen := len(test.ExpectedResult), len(resultSlice); expectedResultLen != resultSliceLen {
				t.Errorf("len(test.ExpectedResult) != len(resultSlice) (%v != %v)", expectedResultLen, resultSliceLen)
			}

			for i := range test.ExpectedResult {
				if test.ExpectedResult[i] != resultSlice[i] {
					t.Errorf(
						"test.ExpectedResult[i] != resultSlice[i] (i = %v, %v != %v)",
						i,
						test.ExpectedResult[i],
						resultSlice[i],
					)
				}
			}

			t.Log(resultSlice, len(resultSlice))

		})
	}

}
