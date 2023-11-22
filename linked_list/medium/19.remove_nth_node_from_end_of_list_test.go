package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/

// method 1 two pointers (fast and slow)
// 1) use two pointers, fast and slow
// 2) move the fast pointer n steps forward
// 3) move the fast pointer to the end of the list, move the slow pointer to the node before the node to be removed
// 4) remove the target node (slow.Next)
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func removeNthFromEnd1(head *util.ListNode, n int) *util.ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	dummy := &util.ListNode{Next: head}

	// use two pointers, fast and slow
	slow := dummy
	fast := dummy

	// move the fast pointer n steps forward
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// move the fast pointer to the end of the list
	// move the slow pointer to the node before the node to be removed
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// remove the target node (slow.Next)
	slow.Next = slow.Next.Next

	return dummy.Next
}

func Test_removeNthFromEnd1(t *testing.T) {
	type args struct {
		head *util.ListNode
		n    int
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
				n:    2,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
				n:    1,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
				n:    1,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
				n:    2,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			removeNthFromEnd1(tc.args.head, tc.args.n),
			tc.name,
		)
	}
}
