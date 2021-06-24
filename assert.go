package testmatrix

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func (matrix TestMatrix) AssertSlice(data [][]interface{}, assert func(subjectText string, headerText string) bool) {
	mreader := NewMatrixReader(data)
	matrix.Assert(mreader, assert)
}

func (matrix TestMatrix) AssertCsv(reader io.Reader, assert func(subjectText string, headerText string) bool) {
	readerCsv := csv.NewReader(reader)
	mreader := NewCsvReader(readerCsv)
	matrix.Assert(mreader, assert)
}

func (matrix TestMatrix) Assert(r Reader, assert func(subjectText string, headerText string) bool) {
	t := matrix.t
	t.Helper()

	matrixSuccess := true
	var buf strings.Builder
	twriter := tablewriter.NewWriter(&buf)
	twriter.SetAutoFormatHeaders(false)
	twriter.SetAlignment(tablewriter.ALIGN_CENTER)

	headers, err := r.ReadHeaders()
	if err != nil {
		t.Fatal(err.Error())
	}

	twriter.SetHeader(headers.Row())

	for {
		results, err := r.Read()

		if err == EOF {
			break
		}
		if err != nil {
			t.Fatal(err.Error())
		}

		outputRow := make([]string, len(headers)+1)
		outputRow[0] = results.Subject

		for headerIndex, header := range headers {
			expected := results.Expected[headerIndex]
			actual := assert(results.Subject, header)

			outputRow[headerIndex+1] = matrix.SuccessText

			if actual != expected {
				outputRow[headerIndex+1] = matrix.FailureText
				matrixSuccess = false
			}
		}

		twriter.Append(outputRow)
	}

	twriter.Render()

	t.Logf("\n%s\n", buf.String())

	if !matrixSuccess {
		t.Error("Matrix test failed")
	}
}
