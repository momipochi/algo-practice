package permutations

import (
	"algo-practice/utils"
	"testing"
)

func TestPermutations(t *testing.T) {
	cases := [][]int{{1, 2, 3}, {0, 1}, {1}}
	ans := [][][]int{{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, {{0, 1}, {1, 0}}, {{1}}}

	for i := 0; i < len(cases); i++ {
		if !utils.IsEqual2DArray(permutations(&cases[i]), ans[i]) {
			t.Fatalf("Permutations, Answer does not match")
		}
	}
}
