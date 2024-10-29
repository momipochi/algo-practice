package l55jumpgame

import (
	"algo-practice/utils"
	"testing"
)

func TestJumpGame(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	out := true
	res := canJump(nums)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
