package easy

import (
	"fmt"
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/reverse-linked-list/

// method 1
// 1) use prev to store the previous node of current node, prev is nil at the beginning
// 2) use cur to store the current node, cur is head at the beginning
// 3) use while loop to scan the list, until cur is nil
// 4) use next to store the next node of current node
// 5) reverse the current node, cur.Next = prev
// 6) move the prev node to the current node, prev = cur
// 7) move the current node to the next node, cur = next
// 8) return prev
func reverseList1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// prev is used to store the previous node of current node
	var prev *util.ListNode
	cur := head

	for cur != nil {
		// store the next node of cur node
		// nil(prev) -> 1(cur) -> 2(cur.Next, next) -> 3 -> 4 -> 5
		next := cur.Next

		// reverse the cur node
		// nil(prev, cur.Next) <- 1(cur), 2(cur.Next, next) -> 3 -> 4 -> 5
		cur.Next = prev

		// move the prev node to the cur node
		// nil(cur.Next) <- 1(cur, prev), 2(next) -> 3 -> 4 -> 5
		prev = cur

		// move the cur node to the next node
		// nil <- 1(prev), 2(cur, next) -> 3(cur.Next) -> 4 -> 5
		cur = next
	}

	return prev
}

func Test_reverseList1(t *testing.T) {
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
				head: util.ArrayToSinglyLinkedList([]int{1, 2, 3, 4, 5}).Head,
			},
			expected: expected{
				result: util.ArrayToSinglyLinkedList([]int{5, 4, 3, 2, 1}).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToSinglyLinkedList([]int{1, 2}).Head,
			},
			expected: expected{
				result: util.ArrayToSinglyLinkedList([]int{2, 1}).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToSinglyLinkedList([]int{}).Head,
			},
			expected: expected{
				result: util.ArrayToSinglyLinkedList([]int{}).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			reverseList1(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
