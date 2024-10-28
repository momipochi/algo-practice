package l39combinationsum

import (
	"algo-practice/utils"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	c := []int{2, 3, 6, 7}
	ans := [][]int{{2, 2, 3}, {7}}
	rec := combinationSum(c, 7)
	if !utils.IsEqual2DArray(rec, ans) {
		utils.DeafultTestOutput(t, ans, rec)
	}
}
