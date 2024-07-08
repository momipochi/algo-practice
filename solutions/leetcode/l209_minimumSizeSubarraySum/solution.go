package l209minimumsizesubarraysum

import "math"

func solution(target *int, nums *[]int) int {
	right, left, sum, ret := 0, 0, 0, math.MaxInt64
	for right < len(*nums) {
		sum += (*nums)[right]
		for sum >= *target {
			ret = min(ret, right-left+1)
			sum -= (*nums)[left]
			left++
		}
		right++
	}
	return ret
}
