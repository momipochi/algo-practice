package l1twosum

import (
	"algo-practice/utils"
	"testing"
)

func Test_twoSum(t *testing.T) {
	nums, want := []int{2, 7, 11, 15}, []int{0, 1}
	target := 9

	if !utils.IsEqualArray(twoSum(nums, target), want) {
		t.Errorf("Wrong answer")
	}
}
