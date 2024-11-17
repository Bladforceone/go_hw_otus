package chessboard

import (
	"errors"
)

func PrintChessboard(size int) ([]string, error) {
	var evenRow, oddRow string

	if size < 2 {
		return nil, errors.New("uncorrected chessboard size")
	}

	for i := 0; i < size; i++ {
		if i%2 == 0 {
			evenRow += " "
			oddRow += "#"
		} else {
			evenRow += "#"
			oddRow += " "
		}
	}

	var ans []string
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			ans = append(ans, evenRow)
		} else {
			ans = append(ans, oddRow)
		}
	}

	return ans, nil
}
