package solutions

//https://leetcode.com/problems/valid-sudoku/

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board[0]); i++ {
		if !isValidSet(&board[i]) {
			return false
		}
	}
	for i := 0; i < len(board[0]); i++ {
		col := []byte{}
		for j := 0; j < len(board[i]); j++ {
			col = append(col, board[j][i])
		}
		if !isValidSet(&col) {
			return false
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			grid := []byte{}
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					grid = append(grid, board[x][y])
				}
			}
			if !isValidSet(&grid) {

				return false
			}
		}
	}
	return true
}

func isValidSet(set *[]byte) bool {
	s := map[byte]bool{}
	for _, b := range *set {
		if b != '.' {
			val, ok := s[b]
			if ok && val {
				return false
			}
			s[b] = true
		}
	}
	return true
}
