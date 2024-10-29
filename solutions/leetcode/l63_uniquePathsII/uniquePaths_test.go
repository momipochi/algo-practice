package l63uniquepathsII

import (
	"algo-practice/utils"
	"testing"
)

func TestUniquePathsII(t *testing.T) {
	grid, out := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2
	res := uniquePathsWithObstacles(grid)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
