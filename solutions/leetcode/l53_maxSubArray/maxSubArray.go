package maxsubarray

//Given an integer array nums, find the subarray with the largest sum, and return its sum.
//https://leetcode.com/problems/maximum-subarray/

func MaxSubArray(nums []int) int {
	cMax, oMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		nextMax := nums[i] + cMax
		cMax = max(nums[i], nextMax)
		oMax = max(cMax, oMax)
	}
	return oMax

}
