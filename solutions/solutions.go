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
