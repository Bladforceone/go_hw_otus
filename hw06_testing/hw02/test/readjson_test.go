package test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw02/reader"
	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw02/types"
	"github.com/stretchr/testify/assert"
)

func TestReadJson(t *testing.T) {
	basePath, _ := os.Getwd()

	tests := []struct {
		name        string
		path        string
		expected    []types.Employee
		expectedErr error
	}{
		{
			name: "MainTest",
			path: filepath.Join(basePath, "data.json"),
			expected: []types.Employee{
				{UserID: 10, Age: 25, Name: "Rob", DepartmentID: 3},
				{UserID: 11, Age: 30, Name: "George", DepartmentID: 2},
			},
			expectedErr: nil,
		},
		{
			name:        "FatalTest_InvalidPath",
			path:        "invalid.json",
			expected:    nil,
			expectedErr: errors.New("open invalid.json: no such file or directory"),
		},
		{
			name:        "FatalTest_InvalidJSON",
			path:        filepath.Join(basePath, "datafatal.json"),
			expected:    nil,
			expectedErr: errors.New("invalid character ':' looking for beginning of object key string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reader.ReadJSON(tt.path)

			if !errors.Is(err, tt.expectedErr) {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expected, got)
		})
	}
}
