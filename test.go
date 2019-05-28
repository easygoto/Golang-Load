package main

import (
	"fmt"
	"unsafe"
)

var f float32 = 3.4
var i, i64, d, cx = 10, 64, 3.099128981278123123, 12 + 3i

const (
	XA = 5 << iota
	XB
	XC
)

func main() {
	dataType()
}

func dataType() {
	var s, ptr = "qwe", &i
	const PI float64 = 3.14
	const NAME string = "tree link"
	const XD = iota
	const XE = iota

	i = 21
	fmt.Println(i, &i)
	fmt.Println(*ptr, ptr)
	fmt.Println(i64+i, d*float64(f), cx*cx, s+"12")

	fmt.Println(PI, len(NAME), unsafe.Sizeof(NAME))

	fmt.Println(XA, XB, XC, XD, XE)
}
