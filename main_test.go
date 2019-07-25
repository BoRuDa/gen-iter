package iter

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	iter := NewIter([]int{1, 3, 2}).
		Reverse().
		Sort().
		ApplyForEach(
			func(p int) int {
				return p + p
			},
			func(p int) int {
				return p * p
			},
		).
		Reverse().
		Reverse()

	i, ok := iter.Next()
	for ok {
		fmt.Println(i, ok)
		i, ok = iter.Next()
	}
}
