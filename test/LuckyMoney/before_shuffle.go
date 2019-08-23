package main

import (
    "../../src/LuckyMoney/infra/algo"
    "fmt"
)

func main() {
    count, amount := int64(10), int64(10)
    for i := int64(0); i < count; i++ {
        x := algo.BeforeShuffle(count, amount*100)
        fmt.Print(x, ",")
    }
    fmt.Println()
}
