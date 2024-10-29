package l79wordsearch

import (
	"algo-practice/utils"
	"testing"
)

func Test_exist(t *testing.T) {
	board, word := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"
	res := exist(board, word)
	if !res {
		utils.DeafultTestOutput(t, true, res)
	}
}
