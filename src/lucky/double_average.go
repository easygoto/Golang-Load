package lucky

import (
	"math/rand"
	"time"
)

// 二倍均值算法
func DoubleAverage(count, amount int64) int64 {
	if count == 1 {
		return amount
	}
	max := amount - min*count
	avg := max / count
	avg2 := 2*avg + min
	rand.Seed(time.Now().UnixNano() + count)
	x := rand.Int63n(avg2) + min
	return x
}
