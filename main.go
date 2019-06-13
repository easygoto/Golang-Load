package main

import "fmt"
import "./HashMap"

func main() {
    testHashMap()
}

func testHashMap() {
    HashMap.GetInstance()
    HashMap.Put("a", "a_put")
    HashMap.Put("b", "b_put")
    fmt.Println(HashMap.Get("a"))
    fmt.Println(HashMap.Get("b"))
    HashMap.Put("p", "p_put")
    fmt.Println(HashMap.Get("p"))
}
