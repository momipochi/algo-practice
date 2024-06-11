package permutations

import (
	"algo-practice/utils"
	"testing"
)

func TestPermutations(t *testing.T) {
	cases := [][]int{{1, 2, 3}, {0, 1}, {1}, {5, 4, 6, 2}}
	ans := [][][]int{{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, {{0, 1}, {1, 0}}, {{1}}, {{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}}}

	for i := 0; i < len(cases); i++ {
		if !utils.IsEqual2DArray(permutations(&cases[i]), ans[i]) {
			t.Fatalf("Permutations, Answer does not match")
		}
	}
}
