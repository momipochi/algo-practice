package l79wordsearch

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])

	var helper func(x, y, bInd int) bool
	helper = func(x, y, bInd int) bool {
		if bInd == len(word) {
			return true
		}
		if x < 0 || y < 0 || x >= row || y >= col || board[x][y] != word[bInd] {
			return false
		}
		tmp := board[x][y]
		board[x][y] = '!'
		res := helper(x+1, y, bInd+1) ||
			helper(x-1, y, bInd+1) ||
			helper(x, y+1, bInd+1) ||
			helper(x, y-1, bInd+1)
		board[x][y] = tmp
		return res
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == word[0] && helper(i, j, 0) {
				return true
			}
		}
	}
	return false
}
