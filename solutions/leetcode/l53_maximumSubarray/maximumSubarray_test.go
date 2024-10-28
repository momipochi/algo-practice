package l53maximumsubarray

import (
	"algo-practice/utils"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	out := 6
	res := maxSubArray(nums)
	if res != out {
		utils.DeafultTestOutput(t, out, res)
	}
}
