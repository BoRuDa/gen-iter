package iter

import (
	"reflect"
	"testing"
)

//go:generate gen-iter -t int -p iter
func TestName(t *testing.T) {
	var (
		expectedResult = []int{16, 36, 4}
		gotResult      = make([]int, 0)
	)

	iter := NewIter([]int{1, 3, 2}).
		Reverse().
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

	val, ok := iter.Next()
	for ok {
		gotResult = append(gotResult, val)
		val, ok = iter.Next()
	}

	if !reflect.DeepEqual(expectedResult, gotResult) {
		t.Fatalf(`expected result: %v, but got: %v`, expectedResult, gotResult)
	}
}
