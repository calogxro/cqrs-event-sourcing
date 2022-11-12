package utils

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(10)
	set.Add(5)
	fmt.Println(set.data)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(10))
	fmt.Println(set.Contains(5))
}
