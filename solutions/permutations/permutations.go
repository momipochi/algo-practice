package permutations

func permutations(arr *[]int) [][]int {
	var res [][]int
	backtrack(&res, arr, []int{}, map[int]struct{}{})
	return res
}

func backtrack(res *[][]int, nums *[]int, permutation []int, used map[int]struct{}) {
	if len(permutation) == len(*nums) {
		p := make([]int, len(permutation))
		copy(p, permutation)
		*res = append(*res, p)
		return
	}
	for i := 0; i < len(*nums); i++ {
		if _, ok := used[i]; !ok {
			used[i] = struct{}{}
			permutation = append(permutation, (*nums)[i])

			backtrack(res, nums, permutation, used)

			delete(used, i)
			permutation = permutation[:len(permutation)-1]
		}
	}
}
