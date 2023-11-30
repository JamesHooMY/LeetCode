package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rotate-list/description/

// method 1
// 1) find the length of the list, and keep the tail
// 2) make the list a cycle
// 3) find the new tail that is the (length - k%length - 1)th node from the head
// 4) the new head is the next of the new tail
// 5) break the cycle of the list from the new tail
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func rotateRight1(head *util.ListNode, k int) *util.ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// find the length of the list, and keep the tail
	length := 1
	tail := head
	for tail.Next != nil {
		length++
		tail = tail.Next
	}

	// make the list a cycle
	tail.Next = head

	// find the new tail
	newTail := head
	for i := 0; i < length-(k%length)-1; i++ {
		newTail = newTail.Next
	}

	// the new head is the next of the new tail
	newHead := newTail.Next

	// break the cycle of the list from the new tail
	newTail.Next = nil

	return newHead
}

func Test_rotateRight1(t *testing.T) {
	type args struct {
		head *util.ListNode
		k    int
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
				k:    2,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{4, 5, 1, 2, 3}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{0, 1, 2}, -1).Head,
				k:    4,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 0, 1}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
				k:    1,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 1}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			rotateRight1(tc.args.head, tc.args.k),
			tc.name,
		)
	}
}
