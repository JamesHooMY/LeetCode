package util_test

import (
	"testing"

	"leetcode/tree/util"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayToBinaryTree(t *testing.T) {
	type args struct {
		arr []int
	}
	type expected struct {
		result *util.TreeNode[int]
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
				arr: []int{4, 2, 7, 1, 3, 6, 9},
			},
			expected: expected{
				result: &util.TreeNode[int]{
					Val: 4,
					Left: &util.TreeNode[int]{
						Val: 2,
						Left: &util.TreeNode[int]{
							Val: 1,
						},
						Right: &util.TreeNode[int]{
							Val: 3,
						},
					},
					Right: &util.TreeNode[int]{
						Val: 7,
						Left: &util.TreeNode[int]{
							Val: 6,
						},
						Right: &util.TreeNode[int]{
							Val: 9,
						},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				arr: []int{1, 2},
			},
			expected: expected{
				result: &util.TreeNode[int]{
					Val: 1,
					Left: &util.TreeNode[int]{
						Val: 2,
					},
				},
			},
		},
		{
			name: "3",
			args: args{
				arr: []int{1, -1, 2},
			},
			expected: expected{
				result: &util.TreeNode[int]{
					Val: 1,
					Right: &util.TreeNode[int]{
						Val: 2,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(
				t,
				tc.expected.result,
				util.ArrayToBinaryTree(tc.args.arr),
				tc.name,
			)
		})
	}
}
