package util_test

import (
	"fmt"
	"testing"

	"leetcode/linked_list/util"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayToCycleOrSinglyLinkedList_Singly(t *testing.T) {
	type args struct {
		arr []int
		pos int
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
				arr: []int{1, 2, 3, 4, 5},
				pos: -1,
			},
			expected: expected{
				result: &util.ListNode{
					Val: 1,
					Next: &util.ListNode{
						Val: 2,
						Next: &util.ListNode{
							Val: 3,
							Next: &util.ListNode{
								Val: 4,
								Next: &util.ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{1},
				pos: -1,
			},
			expected: expected{
				result: &util.ListNode{
					Val: 1,
				},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{},
				pos: -1,
			},
			expected: expected{
				result: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				util.ArrayToCycleOrSinglyLinkedList(tc.args.arr, tc.args.pos).Head,
				fmt.Sprintf("testCase name: %s", tc.name),
			)
		})
	}
}

func Test_ArrayToCycleOrSinglyLinkedList_Cycle(t *testing.T) {
	node1 := &util.ListNode{Val: 1, Next: nil}
	node2 := &util.ListNode{Val: 2, Next: nil}
	node3 := &util.ListNode{Val: 3, Next: nil}
	node4 := &util.ListNode{Val: 4, Next: nil}

	node1.Next = node2
	node2.Next = node3 // cycle start
	node3.Next = node4
	node4.Next = node2 // cycle

	assert.Equal(
		t,
		node1,
		util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, 1).Head,
	)
}

func Test_ArrayToCycleOrSinglyLinkedList_Cycle_wrong(t *testing.T) {
	node1 := &util.ListNode{Val: 1, Next: nil}
	node2 := &util.ListNode{Val: 2, Next: nil}
	node3 := &util.ListNode{Val: 3, Next: nil}
	node4 := &util.ListNode{Val: 4, Next: nil}

	node1.Next = node2 // cycle start
	node2.Next = node3
	node3.Next = node4
	node4.Next = node1 // cycle

	assert.NotEqual(
		t,
		node1,
		util.ArrayToCycleOrSinglyLinkedList([]int{1, 2, 3, 4}, 1).Head,
	)
}
