package oo

type Square struct {
	Width float64
	Rect
}

func (square Square) Area() float64 {
	return square.Width * square.Width
}

func (square Square) GetEdge() int {
	return square.edge
}

func NewSquare(width float64) *Square {
	return &Square{
		Width: width,
		Rect: Rect{
			Height: width,
			edge:   4,
		},
	}
}
