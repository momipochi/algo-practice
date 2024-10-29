package l78subsets

func subsets(nums []int) [][]int {
	res := [][]int{}
	res = append(res, []int{})
	for i := 0; i < len(nums); i++ {
		size := len(res)
		for j := 0; j < size; j++ {
			tmp := make([]int, len(res[j]))
			copy(tmp, res[j])
			tmp = append(tmp, nums[i])
			res = append(res, tmp)
		}
	}
	return res
}
