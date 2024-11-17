package test

import (
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw04/comparator"
	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw04/types"
	"github.com/stretchr/testify/assert"
)

func TestComparator(t *testing.T) {
	BookRate := types.Book{}
	BookRate.SetRate(2.5)
	BookRateBig := types.Book{}
	BookRateBig.SetRate(5)

	BookYear := types.Book{}
	BookYear.SetYear(1999)
	BookYearBig := types.Book{}
	BookYearBig.SetYear(2020)

	BookSize := types.Book{}
	BookSize.SetSize(10)
	BookSizeBig := types.Book{}
	BookSizeBig.SetSize(100)

	tests := []struct {
		name           string
		book1          types.Book
		book2          types.Book
		modeComparator comparator.CompareMode
		expected       bool
		expectedErr    error
	}{
		{name: "CompareRateTrue", book1: BookRateBig, book2: BookRate, modeComparator: comparator.CompareByRate,
			expected: true, expectedErr: nil},
		{name: "CompareYearTrue", book1: BookYearBig, book2: BookYear, modeComparator: comparator.CompareByYear,
			expected: true, expectedErr: nil},
		{name: "CompareSizeTrue", book1: BookSizeBig, book2: BookSize, modeComparator: comparator.CompareBySize,
			expected: true, expectedErr: nil},
		{name: "CompareTestFalse", book1: BookRate, book2: BookRateBig, modeComparator: comparator.CompareByRate,
			expected: false, expectedErr: nil},
		{name: "Equal", book1: BookRate, book2: BookRate, modeComparator: comparator.CompareByRate,
			expected: false, expectedErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comp := comparator.NewComparator(tt.modeComparator)
			got, err := comp.Compare(tt.book1, tt.book2)
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
