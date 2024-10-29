package l78subsets

import (
	"algo-practice/utils"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums, out := []int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
	res := subsets(nums)
	if !utils.IsEqual2DArray(res, out) {
		utils.DeafultTestOutput(t, out, res)
	}
}
