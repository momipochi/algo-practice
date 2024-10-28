package l39combinationsum

// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	combinations := []int{}
	var dfs func(sum int, index int)
	dfs = func(sum int, index int) {
		if sum == target {
			cpy := make([]int, len(combinations))
			copy(cpy, combinations)
			res = append(res, cpy)
			return
		}
		if index == len(candidates) || sum > target {
			return
		}
		combinations = append(combinations, candidates[index])
		dfs(sum+candidates[index], index)
		combinations = combinations[:len(combinations)-1]
		dfs(sum, index+1)
	}
	dfs(0, 0)
	return res
}
