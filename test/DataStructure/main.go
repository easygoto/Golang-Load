package main

import (
	"fmt"

	"DataStructure"
)

func main() {
	testHashMap()
}

func testHashMap() {
	DataStructure.GetInstance()
	DataStructure.Put("a", "a_put")
	DataStructure.Put("b", "b_put")
	fmt.Println(DataStructure.Get("a"))
	fmt.Println(DataStructure.Get("b"))
	DataStructure.Put("p", "p_put")
	fmt.Println(DataStructure.Get("p"))
}
