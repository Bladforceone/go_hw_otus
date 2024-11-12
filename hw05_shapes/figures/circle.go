package figures

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) CalculateArea() float64 {
	s := math.Pow(c.Radius, 2) * math.Pi
	return s
}
