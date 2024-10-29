package l64minimumpathsum

import "math"

func minPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	var helper func(x, y int) int
	helper = func(x, y int) int {
		if x == 0 && y == 0 {
			return grid[x][y]
		}
		if x < 0 || y < 0 {
			return math.MaxInt64
		}
		if dp[x][y] != 0 {
			return dp[x][y]
		}

		dp[x][y] = grid[x][y] + min(helper(x-1, y), helper(x, y-1))
		return dp[x][y]
	}
	return helper(row-1, col-1)
}
