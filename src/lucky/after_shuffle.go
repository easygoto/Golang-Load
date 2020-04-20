package lucky

import "math/rand"

func AfterShuffle(count, amount int64) []int64 {
	ids := make([]int64, 0)

	max := amount - min*count

	remain := max

	for i := int64(0); i < count; i++ {
		x := SimpleRand(count-i, remain)
		remain -= x
		ids = append(ids, x)
	}

	rand.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	return ids
}
