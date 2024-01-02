package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/spiral-matrix/description/

// method 1 iterative
// TC: O(M*N), SC: O(1), M is the number of rows, N is the number of columns
// * this is the best solution for me currently
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	result := []int{}
	rowNum, colNum := len(matrix), len(matrix[0])
	topIndex, bottomIndex, leftIndex, rightIndex := 0, rowNum-1, 0, colNum-1

	// matrix[rowIndex][colIndex]
	for topIndex <= bottomIndex && leftIndex <= rightIndex {
		// left -> right, traverse current topmost row, from leftIndex to rightIndex
		for i := leftIndex; i <= rightIndex; i++ {
			result = append(result, matrix[topIndex][i])
		}
		topIndex++

		// top -> bottom, traverse current rightmost column, from topIndex to bottomIndex
		for i := topIndex; i <= bottomIndex; i++ {
			result = append(result, matrix[i][rightIndex])
		}
		rightIndex--

		// right -> left, traverse current bottommost row, from rightIndex to leftIndex
		if topIndex <= bottomIndex { // this condition make sure we don't traverse the same row again
			for i := rightIndex; i >= leftIndex; i-- {
				result = append(result, matrix[bottomIndex][i])
			}
			bottomIndex--
		}

		// bottom -> top, traverse current leftmost column, from bottomIndex to topIndex
		if leftIndex <= rightIndex { // this condition make sure we don't traverse the same column again
			for i := bottomIndex; i >= topIndex; i-- {
				result = append(result, matrix[i][leftIndex])
			}
			leftIndex++
		}
	}

	return result
}

func Test_spiralOrder1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	type expected struct {
		result []int
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
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			expected: expected{
				result: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			},
			expected: expected{
				result: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			spiralOrder1(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
