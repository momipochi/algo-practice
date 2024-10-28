package l49groupanagrams

import (
	"slices"
)

// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, v := range strs {
		tmp := []byte(v)
		slices.Sort(tmp)
		r := string(tmp)
		m[r] = append(m[r], v)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
