package l77combinations

import (
	"algo-practice/utils"
	"testing"
)

func Test_combine(t *testing.T) {
	n, k, out := 4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	res := combine(n, k)
	if !utils.IsEqual2DArray(res, out) {
		utils.DeafultTestOutput(t, out, res)
	}
}
