package algo

import (
    "math/rand"
    "time"
)

const min = int64(1)

func SimpleRand(count, amount int64) int64 {
    if count == 1 {
        return amount
    }
    max := amount - min*(count-1)
    rand.Seed(time.Now().Unix())
    x := rand.Int63n(max) + min
    return x
}
