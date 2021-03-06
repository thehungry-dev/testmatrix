// Package testmatrix helps running a matrix of tests
package testmatrix

import "testing"

type TestMatrix struct {
	t           *testing.T
	FailureText string
	SuccessText string
}

func NewColoredTestMatrix(t *testing.T) TestMatrix {
	return TestMatrix{t: t, FailureText: "\033[31mâ\033[0m", SuccessText: "\033[32mâ\033[0m"}
}
func NewTestMatrix(t *testing.T) TestMatrix {
	return TestMatrix{t: t, FailureText: "â", SuccessText: "â"}
}

func StringToBool(text string) (bool, bool) {
	switch text {
	case "TRUE":
		return true, true
	case "FALSE":
		return false, true
	default:
		return false, false
	}
}
