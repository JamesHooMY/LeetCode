package util_test

import (
	"testing"

	. "leetcode/graph/util"

	"github.com/stretchr/testify/assert"
)

func TestArrayToGraphList(t *testing.T) {
	type args struct {
		arr [][]int
	}
	type expected struct {
		result *GraphList
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
				arr: [][]int{
					{2, 3},
					{1, 3},
					{1, 2},
				},
			},
			expected: expected{
				result: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 2, Neighbors: []*Node[int]{}},
						}},
					},
				},
			},
		},
		{
			name: "2",
			args: args{
				arr: [][]int{
					{2, 4},
					{1, 3},
					{2, 4},
					{1, 3},
				},
			},
			expected: expected{
				result: &GraphList{
					Vertices: 4,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 4, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 4, Neighbors: []*Node[int]{}},
						}},
						{Val: 4, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
					},
				},
			},
		},
		{
			name: "3",
			args: args{
				arr: [][]int{},
			},
			expected: expected{
				result: &GraphList{
					Vertices: 1,
					Nodes:    []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{}},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ArrayToGraphList(tc.args.arr)
			assert.Equal(t, tc.expected.result.Vertices, result.Vertices)
		})
	}
}

func DeepEqualGraphLists(gl1, gl2 *GraphList) bool {
	if gl1.Vertices != gl2.Vertices {
		return false
	}
	for i := range gl1.Nodes {
		if !DeepEqualNodes(gl1.Nodes[i], gl2.Nodes[i]) {
			return false
		}
	}
	return true
}

func DeepEqualNodes(n1, n2 *Node[int]) bool {
	if n1.Val != n2.Val {
		return false
	}
	if len(n1.Neighbors) != len(n2.Neighbors) {
		return false
	}
	for i := range n1.Neighbors {
		if n1.Neighbors[i].Val != n2.Neighbors[i].Val {
			return false
		}
	}
	return true
}

func TestDeepEqualGraphLists(t *testing.T) {
	type args struct {
		gl1, gl2 *GraphList
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
				gl1: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 2, Neighbors: []*Node[int]{}},
						}},
					},
				},
				gl2: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 2, Neighbors: []*Node[int]{}},
						}},
					},
				},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				gl1: &GraphList{
					Vertices: 2,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
						}},
					},
				},
				gl2: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
						}},
					},
				},
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				gl1: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
						}},
					},
				},
				gl2: &GraphList{
					Vertices: 3,
					Nodes: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{
							{Val: 2, Neighbors: []*Node[int]{}},
							{Val: 3, Neighbors: []*Node[int]{}},
						}},
						{Val: 2, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
						}},
						{Val: 3, Neighbors: []*Node[int]{
							{Val: 1, Neighbors: []*Node[int]{}},
						}},
					},
				},
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DeepEqualGraphLists(tc.args.gl1, tc.args.gl2)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}

func TestDeepEqualNodes(t *testing.T) {
	type args struct {
		n1, n2 *Node[int]
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
				n1: &Node[int]{
					Val: 1,
					Neighbors: []*Node[int]{
						{Val: 2, Neighbors: []*Node[int]{}},
					},
				},
				n2: &Node[int]{
					Val: 1,
					Neighbors: []*Node[int]{
						{Val: 2, Neighbors: []*Node[int]{}},
					},
				},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				n1: &Node[int]{
					Val: 1,
					Neighbors: []*Node[int]{
						{Val: 2, Neighbors: []*Node[int]{}},
					},
				},
				n2: &Node[int]{
					Val: 2,
					Neighbors: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{}},
					},
				},
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				n1: &Node[int]{
					Val: 1,
					Neighbors: []*Node[int]{
						{Val: 2, Neighbors: []*Node[int]{}},
					},
				},
				n2: &Node[int]{
					Val: 2,
					Neighbors: []*Node[int]{
						{Val: 1, Neighbors: []*Node[int]{}},
						{Val: 3, Neighbors: []*Node[int]{}},
					},
				},
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := DeepEqualNodes(tc.args.n1, tc.args.n2)
			assert.Equal(t, tc.expected.result, result)
		})
	}
}
