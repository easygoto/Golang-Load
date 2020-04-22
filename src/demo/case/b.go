package _case

var index = 0

type Bee struct {
	Id int
}

func Bee_(id int) *Bee {
	return &Bee{Id: id}
}

func BeeId(id int) *Bee {
	return &Bee{Id: id}
}

func (Bee) Test() string {
	return "bee"
}

func (bee Bee) SetIndex(in int) *Bee {
	index = in
	return &bee
}

func (Bee) GetIndex() int {
	return index
}

type Book struct {
}

func (Book) Test() string {
	return "book"
}

func (Book) Demo() string {
	return "book demo ..."
}
