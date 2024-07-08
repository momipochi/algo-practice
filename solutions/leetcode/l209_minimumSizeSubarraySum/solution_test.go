package l209minimumsizesubarraysum

import "testing"

func TestMain(t *testing.T) {
	kase, kaseNum := []int{2, 3, 1, 2, 4, 3}, 7
	if solution(&kaseNum, &kase) != 2 {
		t.Fatalf("Minimums size subarray sum, does not match test answer")
	}
}
