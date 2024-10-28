package l50pow

import (
	"algo-practice/utils"
	"testing"
)

func TestMyPow(t *testing.T) {
	x, n, out := 2.0, 10, 1024.0
	res := myPow(x, n)
	if out != res {
		utils.DeafultTestOutput(t, out, res)
	}
}
