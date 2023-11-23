package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/add-two-numbers/description/

// method 1 one pointer (current)
// 1) use a dummy node to store the head of the list
// 2) use a current node to store the node before the next pair
// 3) use while loop to swap the next pair
// 4) move to the node before the next pair
// 5) return the next node of the dummy node
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func addTwoNumbers1(l1 *util.ListNode, l2 *util.ListNode) *util.ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	dummy := &util.ListNode{}
	current := dummy
	carry := 0 // store the carry for the next digit sum

	// carry > 0 means there is a carry for the next digit sum
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		// calculate the sum of the current digit
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		// calculate the carry for the next digit sum
		carry = sum / 10

		// calculate the value of the current digit
		current.Next = &util.ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

func Test_addTwoNumbers1(t *testing.T) {
	type args struct {
		l1 *util.ListNode
		l2 *util.ListNode
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
				l1: util.ArrayToCycleOrSinglyLinkedList([]int{2, 4, 3}, -1).Head,
				l2: util.ArrayToCycleOrSinglyLinkedList([]int{5, 6, 4}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{7, 0, 8}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				l1: util.ArrayToCycleOrSinglyLinkedList([]int{0}, -1).Head,
				l2: util.ArrayToCycleOrSinglyLinkedList([]int{0}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{0}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				l1: util.ArrayToCycleOrSinglyLinkedList([]int{9, 9, 9, 9, 9, 9, 9}, -1).Head,
				l2: util.ArrayToCycleOrSinglyLinkedList([]int{9, 9, 9, 9}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			addTwoNumbers1(tc.args.l1, tc.args.l2),
			tc.name,
		)
	}
}
