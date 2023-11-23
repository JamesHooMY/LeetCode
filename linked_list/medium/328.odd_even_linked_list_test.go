package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/odd-even-linked-list/description/

// method 1 two pointers (odd and even)
// 1) use two pointers, odd and even
// 2) use while loop to move the odd pointer and even pointer
// 3) connect the even list to the odd list
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func oddEvenList1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	evenHead := even // store the head of the even list，used to connect even list to odd list later

	// 1(head, odd) -> 2(evenHead, even) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
	for even != nil && even.Next != nil {
		// 1(head, odd) -> 2(evenHead, even) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
		// 1(head, odd) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8; 2(evenHead, even) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
		// 1(head) -> 3(odd) -> 4 -> 5 -> 6 -> 7 -> 8; 2(evenHead, even) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
		odd.Next = even.Next
		odd = odd.Next

		// 1(head) -> 3(odd) -> 4(odd.Next) -> 5 -> 6 -> 7 -> 8; 2(evenHead, even) -> 3 -> 4 -> 5 -> 6 -> 7 -> 8
		// 1(head) -> 3(odd) -> 4(odd.Next, even.Next) -> 5 -> 6 -> 7 -> 8; 2(evenHead, even) -> 4(even.Next) -> 5 -> 6 -> 7 -> 8
		// 1(head) -> 3(odd) -> 4(odd.Next, even) -> 5 -> 6 -> 7 -> 8; 2(evenHead) -> 4(even) -> 5 -> 6 -> 7 -> 8
		even.Next = odd.Next
		even = even.Next
	}

	// 1(head) -> 3 -> 5 -> 7(odd) -> 2(evenHead) -> 4 -> 6 -> 8(even)
	odd.Next = evenHead

	return head
}

func Test_oddEvenList1(t *testing.T) {
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 3, 5, 2, 4}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{2, 1, 3, 5, 6, 4, 7}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{2, 3, 6, 7, 1, 5, 4}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 3, 2}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			oddEvenList1(tc.args.head),
			tc.name,
		)
	}
}
