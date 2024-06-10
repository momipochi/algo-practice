package removenthfromend

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	input := &ListNode{Val: 1}
	head := input
	for i := 0; i < 4; i++ {
		tmp := ListNode{Val: i + 2}
		input.Next = &tmp
		input = input.Next
	}

	res := RemoveNthFromEnd(head, 2)

	for head != nil || res != nil {
		if head.Val != res.Val {
			t.Fatalf("RemoveNthFromEnd test failed, wrong val.")
			break
		}
		if (head.Next != nil) != (res.Next != nil) {
			t.Fatalf("RemoveNthFromEnd test failed, ListNode length is faulty.")
		}
		head = head.Next
		res = res.Next
	}
}
