package solutions

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

func TestMaxSubArray(t *testing.T) {
	cases := [][]int{{5, 4, -1, 7, 8}, {-2, 1, -3, 4, -1, 2, 1, -5, 4}, {1}}
	ans := []int{23, 6, 1}
	for i := 0; i < len(cases); i++ {
		res := MaxSubArray(cases[i])
		if res != ans[i] {
			t.Fatalf("MaxSubArray, Answers do not match")
		}
	}
}

func TestIsValidSudoku(t *testing.T) {
	testCase := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	ans := true
	if IsValidSudoku(testCase) != ans {
		t.Fatalf("IsValidSudoku, Answer does not match")
	}
	testCase = [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	ans = false
	if IsValidSudoku(testCase) != ans {
		t.Fatalf("IsValidSudoku, Answer does not match")
	}
}
