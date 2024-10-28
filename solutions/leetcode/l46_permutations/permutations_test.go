package permutations

import (
	"algo-practice/utils"
	"testing"
)

func TestPermutations(t *testing.T) {
	c := []int{1, 2, 3}
	a := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	ans := permute(c)

	if !utils.IsEqual2DArray(ans, a) {
		utils.DeafultTestOutput(t, a, ans)
	}
}
