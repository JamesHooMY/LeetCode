package medium

import (
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-list/description/

// method 1 merge sort (top-down)
// 1) split the list into two lists, and sort the two lists
// 2) merge the two sorted lists
// TC = O(NlogN), SC = O(logN)
func sortList1(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// split the list into two lists
	middle := split_from_middle_prev(head)

	// sort the two lists
	left := sortList1(head)
	right := sortList1(middle)

	return mergeSort(left, right)
}

// split the list from the node before the middle node, and return the middle node (slow, the first node of the right list)
func split_from_middle_prev(head *util.ListNode) (middle *util.ListNode) {
	if head == nil || head.Next == nil {
		return head
	}

	// find the middle node, and split the list into two lists from the node before the middle node
	// * key point: use prev to store the node before the middle node, and set prev.Next = nil to split the list into two lists
	var prev *util.ListNode
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// split the list into two lists
	prev.Next = nil

	return slow
}

func mergeSort(left *util.ListNode, right *util.ListNode) (head *util.ListNode) {
	// merge the two lists
	// 1) use a dummy node to store the result
	// 2) use a current node to store the current node of the result
	// 3) use while loop to scan the list1 and list2
	// 4) if list1.Val < list2.Val, append list1 to the result, otherwise append list2 to the result
	// 5) move the current node to the next node of the result
	// 6) if list1 is not nil, append the rest of list1 to the result, otherwise append the rest of list2 to the result
	// 7) return the next node of the dummy node
	dummy := &util.ListNode{}
	current := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}

		current = current.Next
	}

	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}

	return dummy.Next
}

// method 2 merge sort (bottom-up)
// 1) use a dummy node to store the result, and set the next node of the dummy node to the head of the list
// 2) use a length variable to store the length of the list
// 3) use a for loop to split the list into two lists, and merge the two lists
// 4) double the length
// 5) return the next node of the dummy node
// TC = O(NlogN), SC = O(1)
// * this is the best solution for me currently
func sortList2(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	length := 0
	current := head

	// get the length of the list
	for current != nil {
		length++
		current = current.Next
	}

	dummy := &util.ListNode{Next: head}
	// example: 4 -> 2 -> 3 -> 1 -> 6 -> 5, length = 6, step = 1, 2, 4
	for step := 1; step < length; step *= 2 {
		prev := dummy
		current := prev.Next

		for current != nil {
			// left stores the start node of the left list
			left := current

			// right stores the start node of the right list
			// * this step split out the left list, store the rest of the list into right
			right := split_by_step(left, step)

			// current stores the rest of the list after right
			// * this step split out the right list, store the rest of the list into current
			current = split_by_step(right, step)

			// update the dummy.Next to the sorted list
			prev.Next = mergeSort(left, right)

			// move prev to the last node of the sorted list
			for prev.Next != nil {
				prev = prev.Next
			}
		}

		/*
			step = 1:
				1st for current != nil:
					4(current, left, dummy.Next, prev.Next) -> 2 -> 3 -> 1 -> 6 -> 5

					split(left, step):
						4(left) -> nil, return 2(right) -> 3 -> 1 -> 6 -> 5

					split(right, step):
						2(right) -> nil, return 3(current) -> 1 -> 6 -> 5

					merge(left, right):
						4(left) -> nil, 2(right) -> nil, return 2(dummy.Next, prev.Next) -> 4 -> nil

					for prev.Next != nil:
						2(dummy.Next, prev) -> 4(prev.Next) -> nil
						2(dummy.Next) -> 4(prev) -> nil(prev.Next)

				2nd for current != nil:
					3(current, left) -> 1 -> 6 -> 5, 2(dummy.Next) -> 4(prev) -> nil(prev.Next)

					split(left, step):
						3(left) -> nil, return 1(right) -> 6 -> 5

					split(right, step):
						1(right) -> nil, return 6(current) -> 5

					merge(left, right):
						3(left) -> nil, 1(right) -> nil, return 1(prev.Next) -> 3 -> nil

					for prev.Next != nil:
						2(dummy.Next) -> 4(prev) -> 1(prev.Next) -> 3 -> nil
						2(dummy.Next) -> 4 -> 1(prev) -> 3(prev.Next) -> nil
						2(dummy.Next) -> 4 -> 1 -> 3(prev) -> nil(prev.Next)

				3rd for current != nil:
					6(current, left) -> 5, 2(dummy.Next) -> 4 -> 1 -> 3(prev) -> nil(prev.Next)

					split(left, step):
						6(left) -> nil, return 5(right) -> nil

					split(right, step):
						5(right) -> nil, return nil(current)

					merge(left, right):
						6(left) -> nil, 5(right) -> nil, return 5(prev.Next) -> 6 -> nil

					for prev.Next != nil:
						2(dummy.Next) -> 4 -> 1 -> 3(prev) -> 5(prev.Next) -> 6 -> nil
						2(dummy.Next) -> 4 -> 1 -> 3 -> 5(prev) -> 6(prev.Next) -> nil
						2(dummy.Next) -> 4 -> 1 -> 3 -> 5 -> 6(prev) -> nil(prev.Next)

			step = 2:
				for current != nil:
					2(current, left) -> 4 -> 1 -> 3 -> 5 -> 6

					split(left, step):
						2(left) -> 4 -> nil, return  1(right) -> 3 -> 5 -> 6

					split(right, step):
						1(right) -> 3 -> nil, return 5(current) -> 6

					merge(left, right):
						2(left) -> 4 -> nil, 1(right) -> 3 -> nil, return 1(dummy.Next, prev.Next) -> 2 -> 3 -> 4 -> nil

					for prev.Next != nil:
						1(dummy.Next, prev) -> 2(prev.Next) -> 3 -> 4 -> nil
						1(dummy.Next) -> 2(prev) -> 3(prev.Next) -> 4 -> nil
						1(dummy.Next) -> 2 -> 3(prev) -> 4(prev.Next) -> nil
						1(dummy.Next) -> 2 -> 3 -> 4(prev) -> nil(prev.Next)

				for current != nil:
					5(current, left) -> 6, 1(dummy.Next) -> 2 -> 3 -> 4(prev) -> nil(prev.Next)

					split(left, step):
						5(left) -> 6 -> nil, return nil(right)

					split(right, step):
						nil(right), return nil(current)

					merge(left, right):
						5(left) -> 6 -> nil, nil(right), return 5(prev.Next) -> 6 -> nil

					for prev.Next != nil:
						1(dummy.Next) -> 2 -> 3 -> 4(prev) -> 5(prev.Next) -> 6 -> nil
						1(dummy.Next) -> 2 -> 3 -> 4 -> 5(prev) -> 6(prev.Next) -> nil
						1(dummy.Next) -> 2 -> 3 -> 4 -> 5 -> 6(prev) -> nil(prev.Next)
		*/
	}

	return dummy.Next
}

// split the list into two lists (left list and right list), and return the right list
func split_by_step(head *util.ListNode, step int) *util.ListNode {
	if head == nil {
		return nil
	}

	left := head
	// find the last node of the left list, which is determined by the step
	for i := 1; left.Next != nil && i < step; i++ {
		left = left.Next
	}

	// right stores the next node of the last node of the left list
	right := left.Next

	// split the list
	left.Next = nil

	return right
}

// method 3 quick sort
// 1) use the last node as the pivot
// 2) partition the list into two lists, and sort the two lists
// 3) merge the two sorted lists
// TC = O(NlogN), SC = O(logN)
func sortList3(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find the end node
	end := head
	for end.Next != nil {
		end = end.Next
	}

	// quick sort
	quickSort(head, end)

	return head
}

func quickSort(head, end *util.ListNode) {
	// head == end means there is only one node in the list
	// head == end.Next means there is a cycle in the list
	if head == nil || end == nil || head == end || head == end.Next {
		return
	}

	// partition the list
	pivotPrev := partition(head, end)
	pivot := pivotPrev.Next

	// * this is the key point, singly linked list not like slice, we must get the node before the pivot to defined the end node of the left list.
	// * Also, we cannot skip the pivot node in right list. Not like slice 912.sort_array.go, we can skip the pivotIndex and start from pivotIndex+1 in right slice.
	quickSort(head, pivotPrev)
	quickSort(pivot, end)
	/*
		* quickSort(pivot.Next, end) will cause error in test case 4

		3(head) -> 4 -> 1(end) -> nil, after partition:
		1(pivotPrev, head) -> 3(pivot) -> 4(pivot.Next, end) -> nil
	*/
}

func partition(head, end *util.ListNode) (pivotPrev *util.ListNode) {
	if head == nil || end == nil || head == end || head == end.Next {
		return head
	}

	// use the last node as the pivot
	pivot := end

	// nodes before pivotTmp are smaller than the pivot， current for iteration
	pivotTmp, current := head, head

	// * pivotPrev stores the node before the pivot for return
	pivotPrev = pivotTmp

	for current != end {
		// 1) keep the nodes before pivotTmp smaller than the pivot
		// 2) keep the pivotTmp and the nodes after pivotTmp larger than or equal to the pivot
		if current.Val < pivot.Val {
			// swap the val of node smaller than the pivot to pivotTmp
			pivotTmp.Val, current.Val = current.Val, pivotTmp.Val

			pivotPrev = pivotTmp
			pivotTmp = pivotTmp.Next
		}

		current = current.Next
	}

	// swap the pivot and the pivotTmp
	pivotTmp.Val, pivot.Val = pivot.Val, pivotTmp.Val

	return pivotPrev
}

// method 4 heap sort
// TODO: waiting for implementation after heap structure learned

func Test_sortList1(t *testing.T) {
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 5, 3, 4, 0}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 0, 3, 4, 5}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 1}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 3, 4}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortList1(tc.args.head),
			tc.name,
		)
	}
}

func Test_sortList2(t *testing.T) {
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 5, 3, 4, 0}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 0, 3, 4, 5}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 1}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 3, 4}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortList2(tc.args.head),
			tc.name,
		)
	}
}

func Test_sortList3(t *testing.T) {
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
				head: util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, -1).Head,
			},
		},
		{
			name: "2",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 5, 3, 4, 0}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{-1, 0, 3, 4, 5}, -1).Head,
			},
		},
		{
			name: "3",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{}, -1).Head,
			},
		},
		{
			name: "4",
			args: args{
				head: util.ArrayToCycleOrSinglyLinkedList([]int{3, 4, 1}, -1).Head,
			},
			expected: expected{
				result: util.ArrayToCycleOrSinglyLinkedList([]int{1, 3, 4}, -1).Head,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			sortList3(tc.args.head),
			tc.name,
		)
	}
}

// benchmark
func Benchmark_sortList1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortList1(util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head)
	}
}

func Benchmark_sortList2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortList2(util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head)
	}
}

func Benchmark_sortList3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortList3(util.ArrayToCycleOrSinglyLinkedList([]int{4, 2, 1, 3}, -1).Head)
	}
}
