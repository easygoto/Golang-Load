package main

import (
    "../../src/LuckyMoney/infra/algo"
    "fmt"
)

func main() {
    count, amount := int64(10), int64(10)
    temp := amount * 100
    for i := int64(0); i < count; i++ {
        x := algo.DoubleRandom(count-i, temp)
        temp -= x
        fmt.Print(x, ",")
    }
    fmt.Println()
}
