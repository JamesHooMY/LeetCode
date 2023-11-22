package easy

import (
	"fmt"
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-linked-list/description/

// method 1 two pointers (fast and slow) find the middle node, then reverse the second half of the list
// 1) use two pointers, fast and slow
// 2) fast pointer moves two steps each time, slow pointer moves one step each time
// 3) if fast pointer reaches nil, then slow pointer is at the middle node
// 4) reverse the second half of the list
// 5) compare the first half and the second half of the list
// 6) if the value is not equal, then return false
// 7) if the value is equal, then move the current node to the next node
// 8) if the next node of the second half is nil, then return true
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func isPalindrome1(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// use two pointers, fast and slow
	fast := head
	slow := head

	// fast pointer moves two steps each time, slow pointer moves one step each time
	// * this is the key point, find the middle node
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the second half of the list
	/*
		case1: odd number of nodes
			1(head) -> 2 -> 3(slow) -> 4 -> 5(fast) -> nil
		case2: even number of nodes
			1(head) -> 2 -> 3(slow) -> 4 -> nil(fast)
	*/
	// * this is the key point, reverse the second half of the list
	reversedSecondHalfHead := reverseList(slow)
	/*
		*nil is initialized as prev

		case1: odd number of nodes
			1) 1(head) -> 2 -> 3(current) -> 4(current.Next, next) -> 5 -> nil
			2) 1(head) -> 2 -> 3(current) -> *nil(prev, current.Next), 4(next) -> 5 -> nil
			3) 1(head) -> 2 -> 3(current, prev) -> *nil(current.Next), 4(next) -> 5 -> nil
			4) 1(head) -> 2 -> 3(prev) -> *nil, 4(next, current) -> 5(current.Next) -> nil

			1) 1(head) -> 2 -> 3(prev) -> *nil, 4(current) -> 5(current.Next, next) -> nil
			2) 1(head) -> 2 -> 3(prev, current.Next) -> *nil, 4(current) -> 3(prev, current.Next) -> *nil, 5(next) -> nil
			3) 1(head) -> 2 -> 3(current.Next) -> *nil, 4(current, prev) -> 3(current.Next) -> *nil, 5(next) -> nil
			4) 1(head) -> 2 -> 3 -> *nil, 4(prev) -> 3 -> *nil, 5(next, current) -> nil(current.Next)

			1) 1(head) -> 2 -> 3 -> *nil, 4(prev) -> 3 -> *nil, 5(current) -> nil(current.Next, next)
			2) 1(head) -> 2 -> 3 -> *nil, 5(current) -> 4(prev, current.Next) -> 3 -> *nil, nil(next)
			3) 1(head) -> 2 -> 3 -> *nil, 5(current, prev) -> 4(current.Next) -> 3 -> *nil, nil(next)
			4) 1(head) -> 2 -> 3 -> *nil, 5(prev) -> 4 -> 3 -> *nil, nil(next, current)

			finally:
				1(head) -> 2 -> 3(same node) -> *nil(same node)
				5(prev) -> 4 -> 3(same node) -> *nil(same node)

			reversedSecondHalfHead result:
				5(prev) -> 4 -> 3 -> *nil

		case2: even number of nodes
			1) 1(head) -> 2 -> 3(current) -> 4(current.Next, next) -> nil
			2) 1(head) -> 2 -> 3(current) -> *nil(prev, current.Next), 4(next) -> nil
			3) 1(head) -> 2 -> 3(current, prev) -> *nil(current.Next), 4(next) -> nil
			4) 1(head) -> 2 -> 3(prev) -> *nil, 4(next, current) -> nil(current.Next)

			1) 1(head) -> 2 -> 3(prev) -> *nil, 4(current) -> nil(current.Next, next)
			2) 1(head) -> 2 -> 3(prev, current.Next) -> *nil, 4(current) -> 3(prev, current.Next) -> *nil, nil(next)
			3) 1(head) -> 2 -> 3(current.Next) -> *nil, 4(current, prev) -> 3(current.Next) -> *nil, nil(next)
			4) 1(head) -> 2 -> 3 -> *nil, 4(prev) -> 3 -> *nil, nil(next, current)

			finally:
				1(head) -> 2 -> 3(same node) -> *nil(same node)
				4(prev) -> 3(same node) -> *nil(same node)

			reversedSecondHalfHead result:
				4(prev) -> 3 -> *nil
	*/

	// compare the first half and the second half of the list
	for head != nil && reversedSecondHalfHead != nil {
		// if the value is not equal, then return false
		if head.Val != reversedSecondHalfHead.Val {
			return false
		}

		head = head.Next
		reversedSecondHalfHead = reversedSecondHalfHead.Next
	}

	return true
}

func reverseList(head *util.ListNode) *util.ListNode {
	var prev *util.ListNode // *nil
	current := head

	for current != nil {
		next := current.Next // 1)
		current.Next = prev  // 2)
		prev = current       // 3)
		current = next       // 4)
	}

	return prev
}

// method 2 array two pointers
// 1) use an array to store the value of the node
// 2) push the value of the node to the array
// 3) pop the value of the node from the array, compare the value of the current node and the popped value
// 4) if the value is not equal, then return false
// 5) if the value is equal, then move the current node to the next node
// 6) if the next node of the popped value is nil, then return true
// TC = O(N), SC = O(N)
func isPalindrome2(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// use an array to store the value of the node
	arr := make([]int, 0)
	current := head

	// TC = O(N), SC = O(N)
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}

	current = head

	// TC = O(N), SC = O(1)
	for len(arr) > 0 {
		// pop the last value from the array
		lastVal := arr[len(arr)-1]
		arr = arr[:len(arr)-1]

		// compare the value of the current node and the last value
		if current.Val != lastVal {
			return false
		}

		current = current.Next

		// if the next node of last value is nil, then return false
		if current == nil {
			return true
		}
	}

	return true
}

// method 3 stack
// 1) use a stack to store the node
// 2) push the node to the stack
// 3) pop the node from the stack, compare the value of the current node and the popped node
// 4) if the value is not equal, then return false
// 5) if the value is equal, then move the current node to the next node
// 6) if the next node of the popped node is nil, then return true
// TC = O(N), SC = O(N)
func isPalindrome3(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// use a stack to store the node
	stack := make([]*util.ListNode, 0)
	current := head

	// TC = O(N), SC = O(N)
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	current = head

	// TC = O(N), SC = O(1)
	for len(stack) > 0 {
		// pop the last node from the stack
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// compare the value of the current node and the last node
		if current.Val != lastNode.Val {
			return false
		}

		current = current.Next

		// if the next node of last node is nil, then return false
		if current == nil {
			return true
		}
	}

	return true
}

func Test_isPalindrome1(t *testing.T) {
	type args struct {
		head *util.ListNode
	}
	type expected struct {
		result bool
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isPalindrome1(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		head *util.ListNode
	}
	type expected struct {
		result bool
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isPalindrome2(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isPalindrome3(t *testing.T) {
	type args struct {
		head *util.ListNode
	}
	type expected struct {
		result bool
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2}, -1).Head,
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1}, -1).Head,
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isPalindrome3(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome1(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head)
	}
}

func Benchmark_isPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome2(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head)
	}
}

func Benchmark_isPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindrome3(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 2, 1}, -1).Head)
	}
}