package study

import (
	"fmt"
	"reflect"
	"testing"

	. "demo/oo"
)

func showGraph(graph Graph) {
	_, _ = fmt.Printf("type: %v, graph area: %f\n", reflect.TypeOf(graph).Elem(), graph.Area())
}

func showEdge(edge Edge) {
	_, _ = fmt.Printf("type: %v, edge number: %d\n", reflect.TypeOf(edge).Elem(), edge.GetEdge())
}

func TestOO(t *testing.T) {
	rect1 := NewRect(1.2, 3.4)
	rect2 := &Rect{Width: 3.6, Height: 5.8}
	square1 := NewSquare(5.9)
	square2 := &Square{Width: 9.4}
	circle1 := NewCircle(6.8)
	circle2 := &Circle{Radius: 7.7}

	showEdge(rect1)
	showEdge(rect2)
	showEdge(square1)
	showEdge(square2)
	_, _ = fmt.Println()

	showGraph(rect1)
	showGraph(rect2)
	showGraph(square1)
	showGraph(square2)
	showGraph(circle1)
	showGraph(circle2)
}
