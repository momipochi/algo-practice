package l64minimumpathsum

import (
	"algo-practice/utils"
	"testing"
)

func TestMinmumPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	out := 7
	res := minPathSum(grid)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
