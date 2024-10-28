package l3longestsubstringwithoutrepeat

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	mx, l, r := 0, 0, 1
	for r < len(s) {
		tmp := setNewLeft(&s, &l, &r)
		if tmp != 0 {
			l += tmp
			r++
		} else {
			r++
		}
		tmp = r - l
		mx = max(mx, tmp)
	}
	return mx
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
