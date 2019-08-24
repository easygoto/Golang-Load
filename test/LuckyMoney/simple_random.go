package main

import (
    "LuckyMoney/infra/algo"
    "fmt"
)

func main() {
    amount := int64(1000)
    for count := 3; count > 0; count-- {
        x := algo.SimpleRand(int64(count), amount)
        amount -= x
        fmt.Println(x)
    }
}
