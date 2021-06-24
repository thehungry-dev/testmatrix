package testmatrix_test

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/thehungry-dev/testmatrix"
)

//go:embed atoi.csv
var atoiMatrix string

func TestTestMatrixCsv(t *testing.T) {
	t.Run("Test Matrix CSV atoi", func(t *testing.T) {
		matrix := testmatrix.NewColoredTestMatrix(t)
		reader := strings.NewReader(atoiMatrix)
		matrix.AssertCsv(reader, func(input string, expectedOutput string) bool {
			actualOutput, err := strconv.Atoi(input)

			if expectedOutput == "error" {
				return err != nil
			}

			return strconv.Itoa(actualOutput) == expectedOutput
		})
	})
}
