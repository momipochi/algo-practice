package l45jumpgameii

import (
	"algo-practice/utils"
	"testing"
)

func TestJumpGameII(t *testing.T) {
	in := []int{2, 3, 1, 1, 4}
	out := 2
	res := jump(in)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
