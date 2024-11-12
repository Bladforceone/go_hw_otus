package main

import (
	"errors"
	"fmt"
	"github.com/Bladforceone/go_hw_otus/hw05_shapes/figures"
	"github.com/Bladforceone/go_hw_otus/hw05_shapes/shape"
)

func calculateArea(f any) (float64, error) {
	if figure, ok := f.(shape.Shape); ok {
		s := figure.CalculateArea()
		return s, nil
	} else {
		err := errors.New("переданный объект функции calculateArea, не является фигурой")
		return 0, err
	}
}
func main() {
	// Place your code here.
	c := figures.Circle{Radius: 10}
	nf := "Я не фигура"

	s, err := calculateArea(c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Площадь фигуры = %0.3f \n", s)
	}

	s, err = calculateArea(nf)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Площадь фигуры %0.3f \n", s)
	}
}
