package l3longestsubstringwithoutrepeat

import (
	"algo-practice/utils"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	max, l, r := 0, 0, 1
	for r < len(s) {
		tmp := setNewLeft(&s, &l, &r)
		if tmp != 0 {
			l += tmp
			r++
		} else {
			r++
		}
		tmp = r - l
		max = utils.Max(max, tmp)
	}
	return max
}

func setNewLeft(s *string, l *int, r *int) int {
	count := 0
	for i := *l; i < *r; i++ {
		if (*s)[i] == (*s)[*r] {
			count++
		}
	}
	return count
}
