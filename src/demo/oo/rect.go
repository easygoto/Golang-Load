package oo

type Rect struct {
	Width  float64
	Height float64
	edge   int
}

func (rect Rect) Area() float64 {
	return rect.Width * rect.Height
}

func (rect Rect) GetEdge() int {
	return rect.edge
}

func NewRect(width, height float64) *Rect {
	return &Rect{
		Width:  width,
		Height: height,
		edge:   4,
	}
}
