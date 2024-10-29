package l62uniquepaths

import (
	"algo-practice/utils"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	m, n := 3, 7
	out := 28
	res := uniquePaths(m, n)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
