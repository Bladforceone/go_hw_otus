package hw05

import (
	"errors"

	"github.com/Bladforceone/go_hw_otus/hw06_testing/hw05/shape"
)

func CalculateArea(f any) (float64, error) {
	if figure, ok := f.(shape.Shape); ok {
		s := figure.CalculateArea()
		return s, nil
	}

	err := errors.New("переданный объект функции calculateArea, не является фигурой")
	return 0, err
}
