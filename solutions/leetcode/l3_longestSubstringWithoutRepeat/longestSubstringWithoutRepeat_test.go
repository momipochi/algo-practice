package l3longestsubstringwithoutrepeat

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abcabcbb"
	ans := 3
	res := lengthOfLongestSubstring(s)
	if res != ans {
		t.Errorf("Wrong answer")
	}
}
