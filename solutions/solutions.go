package solutions

//Given the head of a linked list, remove the nth node from the end of the list and return its head.
//https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	mapper := map[int]*ListNode{}
	count := 0
	tmp := head
	for tmp != nil {

		mapper[count] = tmp
		tmp = tmp.Next
		count++
	}
	ind := count - n
	if ind == 0 {
		return head.Next
	}

	target := mapper[ind]
	prev := mapper[ind-1]
	prev.Next = target.Next
	return head
}

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

//https://leetcode.com/problems/valid-sudoku/

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
func IsValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board[0]); i++ {
		if !isValidSet(&board[i]) {
			return false
		}
	}
	for i := 0; i < len(board[0]); i++ {
		col := []byte{}
		for j := 0; j < len(board[i]); j++ {
			col = append(col, board[j][i])
		}
		if !isValidSet(&col) {
			return false
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			grid := []byte{}
			for x := i; x < i+3; x++ {
				for y := j; y < j+3; y++ {
					grid = append(grid, board[x][y])
				}
			}
			if !isValidSet(&grid) {

				return false
			}
		}
	}
	return true
}

func isValidSet(set *[]byte) bool {
	s := map[byte]bool{}
	for _, b := range *set {
		if b != '.' {
			val, ok := s[b]
			if ok && val {
				return false
			}
			s[b] = true
		}
	}
	return true
}
