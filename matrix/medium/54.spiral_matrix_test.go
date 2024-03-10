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
	topIdx, btmIdx, leftIdx, rightIdx := 0, rowNum-1, 0, colNum-1

	// matrix[rowIndex][colIndex]
	for topIdx <= btmIdx && leftIdx <= rightIdx {
		// left -> right, traverse current topmost row, from leftIdx to rightIdx
		for i := leftIdx; i <= rightIdx; i++ {
			result = append(result, matrix[topIdx][i])
		}
		topIdx++

		// top -> bottom, traverse current rightmost column, from topIdx to btmIdx
		for i := topIdx; i <= btmIdx; i++ {
			result = append(result, matrix[i][rightIdx])
		}
		rightIdx--

		// right -> left, traverse current bottommost row, from rightIdx to leftIdx
		if topIdx <= btmIdx { // this condition make sure we don't traverse the same row again
			for i := rightIdx; i >= leftIdx; i-- {
				result = append(result, matrix[btmIdx][i])
			}
			btmIdx--
		}

		// bottom -> top, traverse current leftmost column, from btmIdx to topIdx
		if leftIdx <= rightIdx { // this condition make sure we don't traverse the same column again
			for i := btmIdx; i >= topIdx; i-- {
				result = append(result, matrix[i][leftIdx])
			}
			leftIdx++
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
