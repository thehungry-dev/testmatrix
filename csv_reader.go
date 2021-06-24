package testmatrix

import (
	"encoding/csv"
	"io"
)

type csvReader struct {
	r            *csv.Reader
	firstRowRead bool
}

func NewCsvReader(r *csv.Reader) Reader {
	return &csvReader{r, false}
}

func (reader *csvReader) ReadHeaders() (Headers, error) {
	var headers Headers
	r := reader.r

	if reader.firstRowRead {
		return headers, InvalidPositionErr
	}

	csvRow, err := r.Read()
	if err == io.EOF {
		return headers, EOF
	}
	if err != nil {
		return headers, err
	}
	reader.firstRowRead = true

	headers = make(Headers, len(csvRow)-1)

	for index, header := range csvRow {
		if index == 0 {
			continue
		}

		headers[index-1] = header
	}

	return headers, nil
}

func (reader *csvReader) Read() (ResultsRow, error) {
	results := ResultsRow{}
	r := reader.r

	if !reader.firstRowRead {
		return results, InvalidPositionErr
	}

	csvRow, err := r.Read()
	if err == io.EOF {
		return results, EOF
	}
	if err != nil {
		return results, err
	}

	expecteds := make([]bool, len(csvRow)-1)
	results.Expected = expecteds

	for index, resultText := range csvRow {
		if index == 0 {
			results.Subject = resultText
			continue
		}

		expected, ok := StringToBool(resultText)
		if !ok {
			return results, UnparsableErr
		}

		expecteds[index-1] = expected
	}

	return results, nil
}
