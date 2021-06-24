package testmatrix_test

import (
	_ "embed"
	"strconv"
	"testing"

	"github.com/thehungry-dev/testmatrix"
)

var sliceMatrix [][]interface{}

func TestTestMatrixSlice(t *testing.T) {
	t.Run("Test Matrix Slice atoi", func(t *testing.T) {
		matrix := testmatrix.NewColoredTestMatrix(t)
		matrix.AssertSlice(sliceMatrix, func(input string, expectedOutput string) bool {
			actualOutput, err := strconv.Atoi(input)

			if expectedOutput == "error" {
				return err != nil
			}

			return strconv.Itoa(actualOutput) == expectedOutput
		})
	})
}

func init() {
	sliceMatrix = [][]interface{}{
		{"", "10", "1", "error"},
		{"1", "FALSE", "TRUE", "FALSE"},
		{"10", "TRUE", "FALSE", "FALSE"},
		{"a", "FALSE", "FALSE", "TRUE"},
		{"0xa", "FALSE", "FALSE", "TRUE"},
		{"zz", "FALSE", "FALSE", "TRUE"},
	}
}
