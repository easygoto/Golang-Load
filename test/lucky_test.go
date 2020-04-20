package test

import (
	"fmt"
	"lucky"
	"testing"
)

func TestSimpleRand(t *testing.T) {
	ForTest("简单随机算法", t, lucky.SimpleRand)
}

func TestBeforeShuffle(t *testing.T) {
	ForTest("先洗牌算法", t, lucky.BeforeShuffle)
}

// 后洗牌算法
func TestAfterShuffle(t *testing.T) {
	_, _ = fmt.Println(lucky.AfterShuffle(int64(10), int64(10000)))
}

func TestDoubleRandom(t *testing.T) {
	ForTest("二次随机算法", t, lucky.DoubleRandom)
}

func TestDoubleAverage(t *testing.T) {
	ForTest("二倍随机算法", t, lucky.DoubleAverage)
}

func ForTest(message string, t *testing.T, fn func(count, amount int64) int64) {
	count, amount := int64(10), int64(10000)
	remain, sum := amount, int64(0)
	for i := int64(0); i < count; i++ {
		x := fn(count-i, remain)
		remain -= x
		sum += x
		_, _ = fmt.Print(x, ",")
	}
	_, _ = fmt.Println()
}
