package box

var index = 0

type Car struct {
}

func (Car) Test() string {
	return "car"
}

func (Car) GetIndex() int {
	return index
}

type Cat struct {
}

func (Cat) Test() string {
	return "cat"
}
