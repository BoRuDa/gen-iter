package iter

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	ints := []int{1, 2, 3}

	iter := NewIter(ints).
		Reverse().
		Map(func(p int) int {
			return p * p
		}).
		Reverse().
		Reverse()

	i, ok := iter.Next()
	for ok {
		fmt.Println(i, ok)
		i, ok = iter.Next()
	}
}
