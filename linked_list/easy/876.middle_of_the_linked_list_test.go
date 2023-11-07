package easy

import (
	"fmt"
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/middle-of-the-linked-list/

// method 1 two pointers (fast and slow)
// 1) use two pointers, fast and slow
// 2) fast pointer moves two steps each time, slow pointer moves one step each time
// 3) if fast pointer reaches nil, then return slow pointer
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func middleNode1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// use two pointers, fast and slow
	slow := head
	fast := head

	// fast pointer moves two steps each time, slow pointer moves one step each time
	// if fast pointer reaches nil, then return slow pointer
	/*
		case 1: odd number of nodes
		1) 1(slow, fast) -> 2 -> 3 -> 4 -> 5 -> nil
		2) 1 -> 2(slow) -> 3(fast) -> 4 -> 5 -> nil
		3) 1 -> 2 -> 3(slow) -> 4 -> 5(fast) -> nil

		case 2: even number of nodes
		1) 1(slow, fast) -> 2 -> 3 -> 4 -> 5 -> 6 -> nil
		2) 1 -> 2(slow) -> 3(fast) -> 4 -> 5 -> 6 -> nil
		3) 1 -> 2 -> 3(slow) -> 4 -> 5(fast) -> 6 -> nil
		4) 1 -> 2 -> 3 -> 4(slow) -> 5 -> 6(fast) -> nil
	*/
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

// method 2 two while loops, this is more easy to understand
// 1) use a variable to store the length of the list
// 2) use a current node to store the current node of the list
// 3) 1st while loop scan the list and get the length, until current is nil
// 4) 2nd while loop scan the list and get the middle node, until i < length/2
// 5) return the middle node
// TC = O(N + 1/2 N), SC = O(1)
func middleNode2(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	current := head

	for current != nil {
		length++
		current = current.Next
	}

	current = head
	// * this is the key point, we need to use length/2, not length/2+1, because we need to return the middle node, not the next node of the middle node
	for i := 0; i < length/2; i++ {
		current = current.Next
	}

	return current
}

// method 3 array
// 1) use an array to store all the nodes
// 2) return the middle node of the array
// TC = O(N), SC = O(N)
func middleNode3(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	arr := make([]*util.ListNode, 0)
	current := head

	for current != nil {
		arr = append(arr, current)
		current = current.Next
	}

	return arr[len(arr)/2]
}

func Test_middleNode1(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5, 6}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{4, 5, 6}, -1).Head,
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
			middleNode1(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_middleNode2(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5, 6}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{4, 5, 6}, -1).Head,
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
			middleNode2(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_middleNode3(t *testing.T) {
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
				result: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 5}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5, 6}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{4, 5, 6}, -1).Head,
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
			middleNode3(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_middleNode1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		middleNode1(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head)
	}
}

func Benchmark_middleNode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		middleNode2(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head)
	}
}

func Benchmark_middleNode3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		middleNode3(util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4, 5}, -1).Head)
	}
}
