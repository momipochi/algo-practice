package l56mergeintervals

import (
	"algo-practice/utils"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	out := [][]int{{1, 6}, {8, 10}, {15, 18}}
	res := merge(intervals)
	if !utils.IsEqual2DArray(out, res) {
		utils.DeafultTestOutput(t, out, res)
	}
}
