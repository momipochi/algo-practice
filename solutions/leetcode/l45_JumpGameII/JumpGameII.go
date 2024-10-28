package l45jumpgameii

func jump(nums []int) int {
	i, res, len := 0, 0, len(nums)
	var nextIndex func(a, b int) int
	nextIndex = func(a, b int) int {
		res++
		if a+b >= len-1 {
			return len
		}
		max, pos := 0, 0
		for i := a; i <= a+b; i++ {
			tmp := nums[i] + i
			if max < tmp {
				max = tmp
				pos = i
			}
		}
		return pos
	}
	for i < len-1 {
		i = nextIndex(i, nums[i])
	}
	return res
}
