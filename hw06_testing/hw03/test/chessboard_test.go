package test

import (
	"errors"
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw03/chessboard"
	"github.com/stretchr/testify/assert"
)

var ERRSIZE = errors.New("uncorrected chessboard size")

func TestPrintChessboard(t *testing.T) {
	tests := []struct {
		name        string
		size        int
		expected    []string
		expectedErr error
	}{
		{name: "MainTest", size: 5, expected: []string{" # # ", "# # #", " # # ", "# # #", " # # "}, expectedErr: nil},
		{name: "FatalTest", size: 1, expected: nil, expectedErr: ERRSIZE},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chessboard.PrintChessboard(tt.size)
			if tt.expectedErr != nil {
				assert.Error(t, tt.expectedErr, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, got)
			}
		})
	}
}
