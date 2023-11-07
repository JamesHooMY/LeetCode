package easy

import (
	"fmt"
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/linked-list-cycle/

// method 1 hash table
// 1) use a hash table to store the node
// 2) if the node is already in the hash table, then return true
// 3) if the node is not in the hash table, then add the node to the hash table
// 4) if the next node of last node is nil, then return false
// TC = O(N), SC = O(N)
func hasCycle1(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	// use a hash table to store the node
	hashTable := make(map[*util.ListNode]struct{})
	current := head

	for current != nil {
		// if the node is already in the hash table, then return true
		if _, ok := hashTable[current]; ok {
			return true
		}

		// if the node is not in the hash table, then add the node to the hash table
		hashTable[current] = struct{}{}

		// if the next node of last node is nil, then return false
		if current.Next == nil {
			return false
		}

		current = current.Next
	}

	return false
}

// method 2 two pointers (fast and slow)
// 1) use two pointers, fast and slow
// 2) fast pointer moves two steps each time, slow pointer moves one step each time
// 3) if fast pointer meets slow pointer, then return true
// 4) if fast pointer reaches nil, then return false
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func hasCycle2(head *util.ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast := head.Next
	slow := head

	for fast != slow {
		// if fast pointer reaches nil, then return false
		if fast == nil || fast.Next == nil {
			return false
		}

		// fast pointer moves two steps each time, slow pointer moves one step each time
		fast = fast.Next.Next
		slow = slow.Next
	}

	// if fast pointer meets slow pointer, then return true
	return true
}

func Test_hasCycle1(t *testing.T) {
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
				head: util.ArrayToCycleLinkedList([]int{3, 2, 0, -4}).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleLinkedList([]int{1, 2}).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToSinglyLinkedList([]int{1}).Head,
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			hasCycle1(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_hasCycle2(t *testing.T) {
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
				head: util.ArrayToCycleLinkedList([]int{3, 2, 0, -4}).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleLinkedList([]int{1, 2}).Head,
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToSinglyLinkedList([]int{1}).Head,
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			hasCycle2(tc.args.head),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_hasCycle1(b *testing.B) {
	head := util.ArrayToCycleLinkedList([]int{3, 2, 0, -4}).Head

	for i := 0; i < b.N; i++ {
		hasCycle1(head)
	}
}

func Benchmark_hasCycle2(b *testing.B) {
	head := util.ArrayToCycleLinkedList([]int{3, 2, 0, -4}).Head

	for i := 0; i < b.N; i++ {
		hasCycle2(head)
	}
}