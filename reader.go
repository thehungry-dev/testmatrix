package testmatrix

import "errors"

var UnparsableErr error
var InvalidPositionErr error
var EOF error

type Reader interface {
	ReadHeaders() (Headers, error)
	Read() (ResultsRow, error)
}

type Headers []string

func (headers Headers) Row() []string {
	row := make([]string, len(headers)+1)
	row[0] = ""

	for index, header := range headers {
		row[index+1] = header
	}

	return row
}

type ResultsRow struct {
	Subject  string
	Expected []bool
}

func (results ResultsRow) Row(successText string, failureText string) []string {
	row := make([]string, len(results.Expected)+1)
	row[0] = results.Subject

	for index, expected := range results.Expected {
		row[index+1] = failureText
		if expected {
			row[index+1] = successText
		}
	}

	return row
}

func init() {
	UnparsableErr = errors.New("unparsabe value in truth matrix")
	InvalidPositionErr = errors.New("reader position is already after 0")
	EOF = errors.New("testmatrix.EOF")
}
