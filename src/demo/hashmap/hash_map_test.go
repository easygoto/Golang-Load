package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap()
	hashMap.Put("a_test", "a_test")
	hashMap.Put("b_test", "b_test")
	fmt.Println(hashMap.Get("a_test"))
	fmt.Println(hashMap.Get("b_test"))
	hashMap.Put("a_test", "a_test_test")
	hashMap.Put("p_test", "p_test_test")
	fmt.Println(hashMap.Get("a_test"))
	fmt.Println(hashMap.Get("p_test"))
}
