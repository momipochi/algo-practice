package l77combinations

func combine(n int, k int) [][]int {
	res, combinations := [][]int{}, make([]int, 0, k)
	var dfs func(index int)
	dfs = func(index int) {
		size := len(combinations)
		if size == k {
			cpy := make([]int, size)
			copy(cpy, combinations)
			res = append(res, cpy)
			return
		}
		for i := index; i <= n; i++ {
			combinations = append(combinations, i)
			dfs(i + 1)
			combinations = combinations[:len(combinations)-1]
		}
	}
	dfs(1)
	return res
}
