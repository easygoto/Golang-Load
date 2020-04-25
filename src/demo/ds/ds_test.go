package ds_test

import (
	"fmt"
	"testing"

	"demo/ds"
)

func TestHashMap(t *testing.T) {
	hashMap := ds.NewHashMap()
	hashMap.Put("a_test", "a_test")
	hashMap.Put("b_test", "b_test")
	_, _ = fmt.Println(hashMap.Get("a_test"))
	_, _ = fmt.Println(hashMap.Get("b_test"))
	hashMap.Put("a_test", "a_test_test")
	hashMap.Put("p_test", "p_test_test")
	_, _ = fmt.Println(hashMap.Get("a_test"))
	_, _ = fmt.Println(hashMap.Get("p_test"))
}
