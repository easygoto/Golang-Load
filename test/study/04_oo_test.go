package study

import (
	"fmt"
	"testing"

	. "demo/box"
	. "demo/case"
)

func TestOO(t *testing.T) {
	_, _ = fmt.Println(new(Apple).Test())
	_, _ = fmt.Println(new(Bee).Test())
	_, _ = fmt.Println(new(Book).Test())
	_, _ = fmt.Println(new(Book).Demo())
	_, _ = fmt.Println(new(Car).Test())
	_, _ = fmt.Println(new(Cat).Test())

	_, _ = fmt.Println(Bee{Id: 123123}.SetIndex(10).GetIndex())
	_, _ = fmt.Println(new(Bee).GetIndex())
	_, _ = fmt.Println(new(Car).GetIndex())
}
