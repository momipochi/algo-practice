package maxsubarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := [][]int{{5, 4, -1, 7, 8}, {-2, 1, -3, 4, -1, 2, 1, -5, 4}, {1}}
	ans := []int{23, 6, 1}
	for i := 0; i < len(cases); i++ {
		res := MaxSubArray(cases[i])
		if res != ans[i] {
			t.Fatalf("MaxSubArray, Answers do not match")
		}
	}
}
