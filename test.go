package main

import (
	"fmt"
	"time"
	"unsafe"
)

var f float32 = 3.4
var i, i64, d, cx = 10, 64, 3.099128981278123123, 12+3i

const (
	XA = 5 << iota
	XB
	XC
)
const MAX int = 3

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divideR is zero.
    divideE: %d
    divideR: 0
`
	return fmt.Sprintf(strFormat, de.divideE)
}

type DivideError struct {
	divideE int
	divideR int
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

func main() {
	//dataType()
	//control()
	//array()
	//pointArray()
	//pointPoint(&i)
	//_struct()
	//slice()
	//xRange()
	//mapDemo()
	//fmt.Println(fibonacci(15))
	//_interface()

	//result, errorMsg := divide(100, 10)
	//if errorMsg == "" {
	//	fmt.Println("100/10 = ", result)
	//} else {
	//	fmt.Println("errorMsg is: ", errorMsg)
	//}

	go say("--------------------", 5)
	say("++++", 6)
}

func say(s string, total int) {
	for i := 0; i < total; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func divide(varDivideE int, varDivideR int) (result int, errorMsg string) {
	if varDivideR == 0 {
		dData := DivideError{
			divideE: varDivideE,
			divideR: varDivideR,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDivideE / varDivideR, ""
	}

}

func _interface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func mapDemo() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}

func xRange() {
	kvs := map[int]string{1: "apple", 2: "banana"}
	for k, v := range kvs {
		fmt.Printf("%d -> %s\n", k, v)
	}
	for i, c := range "啥" {
		fmt.Println(i, c)
	}
}

func slice() {
	var x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var y = "1234567890"
	fmt.Println(x[:], x[:3], x[3:], x[4:7])
	fmt.Println(y[:], y[:3], y[3:], y[4:7])
}

func _struct() {
	var book1 = Books{"title", "author", "subject", 1}
	fmt.Println(book1)
	fmt.Printf("Book 2 title : %s\n", book1.title)
	fmt.Printf("Book 2 author : %s\n", book1.author)
	fmt.Printf("Book 2 subject : %s\n", book1.subject)
	fmt.Printf("Book 2 bookId : %d\n", book1.bookId)
}

func pointPoint(a *int) {
	var ptr *int
	var pPtr **int

	ptr = a
	pPtr = &ptr

	fmt.Printf("变量 a = %d, &a = %p\n", *a, a)
	fmt.Printf("指针变量 *ptr = %d, ptr = %p\n", *ptr, ptr)
	fmt.Printf("指向指针的指针变量 **pPtr = %d, *pPtr = %p, pPtr = %p\n", **pPtr, *pPtr, pPtr)
}

func pointArray() {
	var a = []int{10, 100, 200}
	var i int
	var ptr [MAX]*int
	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
	}
}

func array() {
	var balance = [10]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance[3] = 50.0
	balance[8] = 12
	fmt.Println(balance)

	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func control() {
	if i < i64 {
		fmt.Println("i < i64")
	} else {
		fmt.Println("i > i64")
	}

	switch i {
	case 1:
		fmt.Println("1")
	case 5:
		fmt.Println("5")
	case 10:
		fmt.Println("10")
		fallthrough
	default:
		fmt.Println("i * i = ", i*i)
		fmt.Println("switch default")
	}

	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Println("x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("received ", i3, " from c3")
		} else {
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no communication")
	}
}

func dataType() {
	var s, ptr = "qwe", &i
	const PI float64 = 3.14
	const NAME string = "tree link"
	const XD = iota
	const XE = iota

	i = 21
	fmt.Println(i, &i, 0.000001+0.00000545744)
	fmt.Println(*ptr, ptr)
	fmt.Println(i64+i, d*float64(f), d/float64(f), cx*cx, s+"12")

	fmt.Println(PI, len(NAME), unsafe.Sizeof(NAME))

	fmt.Println(XA, XB, XC, XD, XE)
}
