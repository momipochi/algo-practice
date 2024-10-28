package permutations

func permute(nums []int) [][]int {
	used := []bool{}
	for i := 0; i < len(nums); i++ {
		used = append(used, false)
	}
	perm := []int{}
	res := [][]int{}
	var dfs func()
	dfs = func() {
		if len(perm) == len(nums) {
			cpy := make([]int, len(perm))
			copy(cpy, perm)
			res = append(res, cpy)
			return
		}
		for i, val := range nums {
			if used[i] {
				continue
			}
			perm = append(perm, val)
			used[i] = true
			dfs()
			perm = perm[:len(perm)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}
