package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reorder-list/description/

// method 1
// 1) find the middle node with slow and fast pointers
// 2) split the list into two lists
// 3) reverse the second half list
// 4) interleave merge the two lists
// TC = O(N), SC = O(1)
func reorderList1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find the middle node with slow and fast pointers
	middle := findMiddle(head)
	secondHalf := middle.Next

	// split the list into two lists
	middle.Next = nil

	// reverse the second half list
	reversedSecondHalfHead := reverseList(secondHalf)

	// interleave merge the two lists
	return interleaveMerge(head, reversedSecondHalfHead)
}

func findMiddle(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverseList(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *util.ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func interleaveMerge(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	dummy := &util.ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		current.Next = l1
		l1 = l1.Next
		current = current.Next

		current.Next = l2
		l2 = l2.Next
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

func Test_reorderList1(t *testing.T) {
	type args struct {
		head *util.ListNode
	}
	type expected struct {
		result *util.ListNode
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 4, 2, 3}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 5, 2, 4, 3}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			reorderList1(tc.args.head),
			tc.name,
		)
	}
}
