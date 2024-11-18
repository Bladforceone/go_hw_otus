package bsearch_test

import (
	"errors"
	"testing"

	"github.com/fixme_my_friend/hw08_binary_search/bsearch"
	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name    string
		massive []int
		target  int
		want    int
		wantErr error
	}{
		{
			name:    "array length less than 3",
			massive: []int{1, 2},
			target:  1,
			want:    -1,
			wantErr: errors.New("uncorrected size"),
		},
		{
			name:    "target found",
			massive: []int{1, 2, 3, 4, 5},
			target:  5,
			want:    4,
			wantErr: nil,
		},
		{
			name:    "target not found",
			massive: []int{1, 2, 3, 4, 5},
			target:  6,
			want:    -1,
			wantErr: errors.New("target not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := bsearch.BinarySearch(test.massive, test.target)
			if !errors.Is(err, test.wantErr) {
				assert.Error(t, err, test.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, test.want, got)
		})
	}
}
