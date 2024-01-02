package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/set-matrix-zeroes/description/

// method 1 in-place solution
// 1) use two variables to store whether first row and first column should be set to 0
// 2) use first row and first column to store which row and column should be set to 0
// 3) set rows and columns to 0
// 4) if first row should be set to 0, set first row to 0
// 5) if first column should be set to 0, set first column to 0
// TC: O(M*N), SC: O(1), M is the number of rows, N is the number of columns
// * this is the best solution for me currently
func setZeroes1(matrix [][]int) [][]int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	// use two variables to store whether first row and first column should be set to 0
	firstRowZero := false
	firstColZero := false

	// check whether first row should be set to 0
	for i := 0; i < colNum; i++ {
		if matrix[0][i] == 0 {
			firstRowZero = true
			break
		}
	}

	// check whether first column should be set to 0
	for i := 0; i < rowNum; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	/*
		use first row and first column to store which row and column should be set to 0

		eg. matrix:
		1 1 1 1	   1 0 1 1
		1 0 1 1 => 0 0 1 1
		1 1 1 1	   1 1 1 1
	*/
	for i := 1; i < rowNum; i++ {
		for j := 1; j < colNum; j++ {
			if matrix[i][j] == 0 {
				// set first row and first column to 0
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	/*
		set rows and columns to 0

		eg. matrix:
		1 0 1 1	   1 0 1 1
		0 0 1 1 => 0 0 0 0
		1 1 1 1	   1 0 1 1
	*/
	for i := 1; i < rowNum; i++ {
		for j := 1; j < colNum; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				// set rows and columns to 0
				matrix[i][j] = 0
			}
		}
	}

	// set first row to 0
	if firstRowZero {
		for i := 0; i < colNum; i++ {
			matrix[0][i] = 0
		}
	}

	// set first column to 0
	if firstColZero {
		for i := 0; i < rowNum; i++ {
			matrix[i][0] = 0
		}
	}

	return matrix
}

// method 2 two arrays
// 1) use two arrays to store which row and column should be set to 0
// 2) set rows and columns to 0
// TC: O(M*N), SC: O(M+N), M is the number of rows, N is the number of columns
func setZeroes2(matrix [][]int) [][]int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	// use two arrays to store which row and column should be set to 0
	rows := make([]bool, rowNum)
	cols := make([]bool, colNum)

	/*
		store which row and column should be set to 0

		eg. matrix:
		1 1 1 1		rows: [false, true, false]
		1 0 1 1 =>
		1 1 1 1		cols: [false, true, false, false]
	*/
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	/*
		set rows and columns to 0

		eg. matrix:
		1 0 1 1
		0 0 0 0
		1 0 1 1
	*/
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

func Test_setZeroes1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	type expected struct {
		result [][]int
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
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0, 0},
					{0, 4, 5, 0},
					{0, 3, 1, 0},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			setZeroes1(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_setZeroes2(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	type expected struct {
		result [][]int
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
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0, 0},
					{0, 4, 5, 0},
					{0, 3, 1, 0},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			setZeroes2(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
