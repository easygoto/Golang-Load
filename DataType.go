package main

import "fmt"

func main() {
	var f float32 = 3.4
	var i, i64, d, cx = 10, 64, 3.099128981278123123, 12 + 3i
	var s, ptr = "qwe", uintptr(i)

	i = 21

	fmt.Println(i64 + i)
	fmt.Println(f)
	fmt.Println(d * float64(f))
	fmt.Println(cx * cx)
	fmt.Println(s + "12")
	fmt.Println(ptr)
}
