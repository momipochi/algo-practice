package l56mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	list, prev := make([][]int, 0, len(intervals)), intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] < intervals[i][0] {
			list = append(list, prev)
			prev = intervals[i]
		} else {
			if intervals[i][1] > prev[1] {
				prev[1] = intervals[i][1]
			}
		}
	}
	list = append(list, prev)
	return list
}
