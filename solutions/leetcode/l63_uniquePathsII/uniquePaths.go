package l63uniquepathsII

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if obstacleGrid[i][j] == 0 {
				if i == 0 && j == 0 {
					dp[i][j] = 1
					continue
				}
				if i > 0 {
					dp[i][j] += dp[i-1][j]
				}
				if j > 0 {
					dp[i][j] += dp[i][j-1]
				}
			}
		}
	}
	return dp[row-1][col-1]
}
