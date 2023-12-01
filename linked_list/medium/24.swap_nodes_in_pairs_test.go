package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/swap-nodes-in-pairs/

// method 1 one pointer (current)
// 1) use a dummy node to store the head of the list
// 2) use a current node to store the node before the next pair
// 3) use while loop to swap the next pair
// 4) move to the node before the next pair
// 5) return the next node of the dummy node
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func swapPairs1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &util.ListNode{Next: head}
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		firstNode := current.Next
		secondNode := current.Next.Next

		// swap
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		current.Next = secondNode

		// move to the node before the next pair
		current = current.Next.Next
	}

	return dummy.Next
}

func Test_swapPairs1(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 1, 4, 3}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			swapPairs1(tc.args.head),
			tc.name,
		)
	}
}
