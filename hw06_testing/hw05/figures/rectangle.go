package figures

type Rectangle struct {
	Height float64
	Width  float64
}

func (c Rectangle) CalculateArea() float64 {
	s := c.Height * c.Width
	return s
}
