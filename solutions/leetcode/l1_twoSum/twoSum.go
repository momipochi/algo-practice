package l1twosum

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		s := target - nums[i]

		if val, ok := mp[s]; ok {
			return []int{val, i}
		}
		mp[nums[i]] = i
	}
	return []int{}
}
