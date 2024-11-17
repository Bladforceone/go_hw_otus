package figures

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) CalculateArea() float64 {
	s := 0.5 * t.Height * t.Base
	return s
}
