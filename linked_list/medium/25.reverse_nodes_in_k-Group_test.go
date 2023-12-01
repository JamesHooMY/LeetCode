package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-nodes-in-k-group/description/

// method 1 two pointers (start, end)
// 1) use a dummy node to store the head of the list, prev to store the node before start
// 2) use start to store the start node for reversing, end to store the end node for reversing
// 3) use while loop to find the end node for reversing, and break the list from the end node for reversing
// 4) reverse the list from the start node, then the start will be the end node of the reversed list, connect it to the next node, which is the next start node for reversing
// 5) move prev to the start, which is the node before the next start node for reversing
// 6) move start and end to the next start node for reversing, for the next round iteration
// 7) return the next node of the dummy node
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func reverseKGroup1(head *util.ListNode, k int) *util.ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	dummy := &util.ListNode{Next: head}

	/*
		prev: store the node before start
		start: store the start node for reversing
		end: store the end node for reversing
	*/
	prev, start, end := dummy, head, head

	for end != nil {
		// find the end node for reversing
		for i := 0; i < k-1 && end != nil; i++ {
			end = end.Next
		}

		// * if end is nil, mean the nodes of list for reversing less than k
		if end == nil {
			break
		}

		// store the next node of the end node for reversing, which is the next start node for reversing
		next := end.Next

		// break the list from the end node for reversing
		end.Next = nil

		// reverse the list from the start node
		prev.Next = reverseList(start)

		// start is the end node of the reversed list, connect it to the next node, which is the next start node for reversing
		start.Next = next

		// adjust the prev to the start, then move start, and end to the next start node for reversing, for the next round iteration
		prev, start, end = start, next, next
	}

	return dummy.Next
}

// method 2 two pointers (start, end), use count to confirm the nodes enough for reversing
// logic thinking is same as method 1, but this logic control is more clear and easy to understand
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func reverseKGroup2(head *util.ListNode, k int) *util.ListNode {
	if head == nil || head.Next == nil || k <= 1 {
		return head
	}

	dummy := &util.ListNode{Next: head}
	prev, start, end := dummy, head, head

	// count: count the nodes for reversing
	count := 0

	for end != nil {
		count++

		// count%k == 0 means the nodes enough for reversing
		if count%k == 0 {
			// the following logic is the same as method 1
			next := end.Next
			end.Next = nil

			prev.Next = reverseList(start)

			start.Next = next
			prev, start, end = start, next, next
		} else {
			// find the end node for reversing
			end = end.Next
		}
	}

	return dummy.Next
}

func Test_reverseKGroup1(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 1, 4, 3, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
				k:    3,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3, 2, 1, 4, 5}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
				k:    1,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1,2,3,4}, -1).Head,
				k:    3,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3,2,1,4}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			reverseKGroup1(tc.args.head, tc.args.k),
			tc.name,
		)
	}
}

func Test_reverseKGroup2(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 1, 4, 3, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
				k:    3,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3, 2, 1, 4, 5}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
				k:    1,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1,2,3,4}, -1).Head,
				k:    3,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3,2,1,4}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			reverseKGroup2(tc.args.head, tc.args.k),
			tc.name,
		)
	}
}
