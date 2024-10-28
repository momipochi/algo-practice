package l49groupanagrams

import (
	"algo-practice/utils"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	out := utils.FlattenArray([][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}})
	ans := utils.FlattenArray(groupAnagrams(in))
	sort.Strings(out)
	sort.Strings(ans)
	if !utils.IsEqualArrayString(ans, out) {
		utils.DeafultTestOutput(t, out, ans)
	}
}
