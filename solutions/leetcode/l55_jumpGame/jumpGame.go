package l55jumpgame

func canJump(nums []int) bool {
	dist := 0
	for i, v := range nums {
		if i > dist {
			return false
		}
		dist = max(dist, i+v)
	}
	return true
}
