package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/valid-sudoku/description/

// method 1 hash map
// TC: O(1), SC: O(1)
// * this is the best solution for me currently
func isValidSudoku1(board [][]byte) bool {
	rowSliceMap := make([]map[byte]bool, 9) // index i means row i, rowSliceMap[i][num] means whether num is in row i
	colSliceMap := make([]map[byte]bool, 9) // index i means col i, colSliceMap[i][num] means whether num is in col i
	boxSliceMap := make([]map[byte]bool, 9) // index i means box i, boxSliceMap[i][num] means whether num is in box i

	// initialize map of slice, make sure map is not nil in slice
	for i := 0; i < 9; i++ {
		rowSliceMap[i] = map[byte]bool{}
		colSliceMap[i] = map[byte]bool{}
		boxSliceMap[i] = map[byte]bool{}
	}

	// iterate board
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}

			// check row
			if _, exist := rowSliceMap[i][num]; exist {
				return false
			} else {
				rowSliceMap[i][num] = true
			}

			// check col
			if _, exist := colSliceMap[j][num]; exist {
				return false
			} else {
				colSliceMap[j][num] = true
			}

			// check box
			boxIndex := i/3*3 + j/3 // boxIndex = (rowIndex/3)*3 + colIndex/3
			if _, exist := boxSliceMap[boxIndex][num]; exist {
				return false
			} else {
				boxSliceMap[boxIndex][num] = true
			}
		}
	}

	return true
}

func Test_isValidSudoku1(t *testing.T) {
	type args struct {
		board [][]byte
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
				board: [][]byte{
					{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				board: [][]byte{
					{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
					{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
					{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
					{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
					{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
					{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
					{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
					{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
					{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
				},
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
			isValidSudoku1(tc.args.board),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
