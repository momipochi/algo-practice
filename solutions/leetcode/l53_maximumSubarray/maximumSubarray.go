package l53maximumsubarray

func maxSubArray(nums []int) int {
	res, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sum = max(nums[i], nums[i]+sum)
		res = max(sum, res)
	}
	return res
}
