package hashmap

import (
	"fmt"
	"testing"
)

func TestHashMap(t *testing.T) {
	GetInstance()
	Put("a", "a_put")
	Put("b", "b_put")
	fmt.Println(Get("a"))
	fmt.Println(Get("b"))
	Put("p", "p_put")
	fmt.Println(Get("p"))
}
