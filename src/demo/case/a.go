package _case

type Thing interface {
	Test() string
}

type Some interface {
	Demo() string
}

type Apple struct {
	Food
	weight float32
}

type Food struct {
	name string
}

func (apple Apple) Test() string {
	return "apple"
}
