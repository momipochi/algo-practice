package l90subsetsii

import (
	"algo-practice/utils"
	"testing"
)

func Test_subsetsWithDup(t *testing.T) {
	nums, out := []int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	res := subsetsWithDup(nums)
	if !utils.IsEqual2DArray(out, res) {
		utils.DeafultTestOutput(t, out, res)
	}
}
