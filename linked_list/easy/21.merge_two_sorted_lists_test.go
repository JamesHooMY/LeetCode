package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-sorted-lists/

// method 1
// 1) use a dummy node to store the result
// 2) use a current node to store the current node of the result
// 3) use while loop to scan the list1 and list2
// 4) if list1.Val < list2.Val, append list1 to the result, otherwise append list2 to the result
// 5) move the current node to the next node of the result
// 6) if list1 is not nil, append the rest of list1 to the result, otherwise append the rest of list2 to the result
// 7) return the next node of the dummy node
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	// use a dummy node to store the result in dummy.Next
	dummy := &ListNode{}
	// use a current node to store the last node of the result
	current := dummy

	// list1 or list2 will be nil after the loop is finished, due to the process of list1 = list1.Next or list2 = list2.Next
	for list1 != nil && list2 != nil {
		// current node is the last node of the result, so we need to append the smaller node to the result through updating the current.Next
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		// move the current node to the next node of the result, keep the current node always the last node of the result
		current = current.Next
	}

	// if list1 is not nil, append the rest of list1 to the result, otherwise append the rest of list2 to the result, because list1 and list2 are sorted
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func Test_mergeTwoLists1(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	type expected struct {
		result *ListNode
	}
	type testCase struct {
		name     string
		args     args
		expected expected
	}

	testCases := []testCase{
		{
			name: "1",
			args: args{
				list1: arrayToLinkedList([]int{1, 2, 4}).Head,
				list2: arrayToLinkedList([]int{1, 3, 4}).Head,
			},
			expected: expected{
				result: arrayToLinkedList([]int{1, 1, 2, 3, 4, 4}).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			mergeTwoLists1(tc.args.list1, tc.args.list2),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

type LinkedList struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func arrayToLinkedList(arr []int) *LinkedList {
	ll := &LinkedList{}
	current := ll.Head

	for _, value := range arr {
		if current == nil {
			ll.Head = &ListNode{Val: value, Next: nil}
			current = ll.Head
		} else {
			current.Next = &ListNode{Val: value, Next: nil}
			current = current.Next
		}
	}

	return ll
}
