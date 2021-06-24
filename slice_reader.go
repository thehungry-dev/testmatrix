package testmatrix

type sliceReader struct {
	matrix   [][]interface{}
	position int
}

func NewMatrixReader(testMatrix [][]interface{}) *sliceReader {
	return &sliceReader{matrix: testMatrix, position: 0}
}

func (reader *sliceReader) ReadHeaders() (Headers, error) {
	matrix := reader.matrix
	var headers Headers

	if len(matrix) < 1 {
		return headers, EOF
	}

	headers = make(Headers, len(matrix[0])-1)

	if reader.position != 0 {
		return headers, InvalidPositionErr
	}

	for index, header := range matrix[0] {
		if index == 0 {
			continue
		}

		textHeader, ok := header.(string)
		if !ok {
			return headers, UnparsableErr
		}

		headers[index-1] = textHeader
	}

	reader.position += 1

	return headers, nil
}

func (reader *sliceReader) Read() (ResultsRow, error) {
	matrix := reader.matrix
	results := ResultsRow{}
	expecteds := make([]bool, len(matrix)-1)
	results.Expected = expecteds

	if reader.position == 0 {
		return results, InvalidPositionErr
	}

	pos := reader.position
	if pos >= len(matrix) {
		return results, EOF
	}

	for index, result := range matrix[pos] {
		if index == 0 {
			subject, ok := result.(string)
			if !ok {
				return results, UnparsableErr
			}

			results.Subject = subject
			continue
		}

		expectedText, ok := result.(string)
		if !ok {
			return results, UnparsableErr
		}

		expected, ok := StringToBool(expectedText)
		if !ok {
			return results, UnparsableErr
		}

		expecteds[index-1] = expected
	}

	reader.position += 1

	return results, nil
}
