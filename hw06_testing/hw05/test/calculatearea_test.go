package test

import (
	"errors"
	"testing"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw05"
	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw05/figures"
	"github.com/stretchr/testify/assert"
)

var ErrNofigure = errors.New("переданный объект функции calculateArea, не является фигурой")

func TestCalculateArea(t *testing.T) {
	Circle := figures.Circle{Radius: 5}
	Rectangle := figures.Rectangle{Height: 5, Width: 5}
	Triangle := figures.Triangle{Height: 5, Base: 5}
	Crash := "No figure"

	tests := []struct {
		name        string
		figure      any
		expected    float64
		expectedErr error
	}{
		{name: "CircleArea", figure: Circle, expected: 78.53981633974483, expectedErr: nil},
		{name: "RectangleArea", figure: Rectangle, expected: 25, expectedErr: nil},
		{name: "TriangleArea", figure: Triangle, expected: 12.5, expectedErr: nil},
		{name: "TestFail", figure: Crash, expected: 0, expectedErr: ErrNofigure},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := hw05.CalculateArea(test.figure)
			if !errors.Is(err, test.expectedErr) {
				assert.Error(t, test.expectedErr, err)
				assert.Equal(t, test.expected, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, got)
			}
		})
	}
}
