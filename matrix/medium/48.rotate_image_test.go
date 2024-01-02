package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rotate-image/description/

// method 1 tmpMatrix
// 1) create tmp matrix
// 2) copy and rotate 90 degree clockwise from matrix to tmp matrix
// 3) copy tmp matrix back to matrix
// TC: O(2*N^2), SC: O(N^2), N is the number of rows or columns
func rotate1(matrix [][]int) [][]int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	// create tmp matrix
	tmpMatrix := make([][]int, rowNum)
	for i := 0; i < rowNum; i++ {
		tmpMatrix[i] = make([]int, colNum)
	}

	/*
		copy and rotate 90 degree clockwise from matrix to tmp matrix

		matrix 90 degree clockwise rotate => tmpMatrix[j][colNum-1-i] = matrix[i][j]
	*/
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			/*
				i is row index, j is col index

				eg. i = 0, j = 1
				tmpMatrix[1][3-1-0] = matrix[0][1] => tmpMatrix[1][2] = matrix[0][1]

				matrix		tmpMatrix
				1 2 3		. . 1
				4 5 6  =>	. . 2
				7 8 9		. . .
			*/
			tmpMatrix[j][colNum-1-i] = matrix[i][j]
		}
	}

	// copy tmp matrix to matrix
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			matrix[i][j] = tmpMatrix[i][j]
		}
	}

	return matrix
}

// method 2 transpose matrix and reverse each row, need to think about why it works
// 1) transpose matrix
// 2) reverse each row
// 3) return matrix
// TC: O(2*N^2), SC: O(1), N is the number of rows or columns
func rotate2(matrix [][]int) [][]int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	// transpose matrix
	/*
		1 2 3		1 4 7
		4 5 6  =>	2 5 8
		7 8 9		3 6 9
	*/
	for i := 0; i < rowNum; i++ {
		for j := i; j < colNum; j++ {
			/*
				i is row index, j is col index

				eg. i = 0, j = 1
				matrix[0][1], matrix[1][0] = matrix[1][0], matrix[0][1] => 2, 4 = 4, 2

				1 2 3		1 4 3					1 4 7
				4 5 6  =>	2 5 6 => transpose => 	2 5 8
				7 8 9		7 8 9					3 6 9
			*/
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse each row with j < colNum/2
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum/2; j++ {
			/*
				i is row index, j is col index

				eg. i = 0, j = 0
				matrix[0][0], matrix[0][3-1-0] = matrix[0][3-1-0], matrix[0][0] =>
				matrix[0][0], matrix[0][2] = matrix[0][2], matrix[0][0] => 1, 7 = 7, 1

				1 4 7	 7 4 1
				2 5 8 => 2 5 8
				3 6 9	 3 6 9

				* eg. i = 0, j = 1, middle column no need to swap, j < colNum/2, colNum/2 = 3/2 = 1
				matrix[0][1], matrix[0][3-1-1] = matrix[0][3-1-1], matrix[0][1] =>
				matrix[0][1], matrix[0][1] = matrix[0][1], matrix[0][1] =>

				7 4 1	 7 4 1
				2 5 8 => 2 5 8
				3 6 9	 3 6 9

				eg. i = 1, j = 0
				matrix[1][0], matrix[1][3-1-0] = matrix[1][3-1-0], matrix[1][0] =>
				matrix[1][0], matrix[1][2] = matrix[1][2], matrix[1][0] => 4, 8 = 8, 4

				7 4 1	 7 4 1
				2 5 8 => 8 5 2
				3 6 9	 3 6 9

				* eg. i = 1, j = 1, middle column no need to swap, j < colNum/2, colNum/2 = 3/2 = 1
				matrix[1][1], matrix[1][3-1-1] = matrix[1][3-1-1], matrix[1][1] =>
				matrix[1][1], matrix[1][1] = matrix[1][1], matrix[1][1] => 5, 5 = 5, 5

				7 4 1	 7 4 1
				8 5 2 => 8 5 2
				3 6 9	 3 6 9

			*/
			matrix[i][j], matrix[i][colNum-1-j] = matrix[i][colNum-1-j], matrix[i][j]
		}
	}

	return matrix
}

// method 3 in-place rotate four rectangles
// 1) rotate four rectangles in-place from outer to inner with i < rowNum/2 && j < (colNum+1)/2
// 2) return matrix
//	TC: O(N^2), SC: O(1), N is the number of rows or columns
// * this is the best solution for me currently
func rotate3(matrix [][]int) [][]int {
	rowNum := len(matrix)
	colNum := len(matrix[0])

	// rotate four rectangles
	for i := 0; i < rowNum/2; i++ {
		for j := 0; j < (colNum+1)/2; j++ {
			/*
				eg. i = 0, j = 0
				matrix[0][0], matrix[0][2], matrix[2][2], matrix[2][0] = matrix[2][0], matrix[0][0], matrix[0][2], matrix[2][2] => 1, 3, 9, 7 = 7, 1, 3, 9

				1 2 3		7 2 1
				4 5 6  =>	4 5 6
				7 8 9		9 8 3

				eg. i = 0, j = 1
				matrix[0][1], matrix[1][2], matrix[2][1], matrix[2][0] = matrix[2][0], matrix[0][1], matrix[1][2], matrix[2][1] => 2, 6, 8, 4 = 4, 2, 6, 8

				7 2 1		7 4 1
				4 5 6  =>	8 5 2
				9 8 3		9 6 3

				eg. i = 1, j = 1
				matrix[1][1], matrix[1][1], matrix[1][1], matrix[1][1] = matrix[1][1], matrix[1][1], matrix[1][1], matrix[1][1] => 5, 5, 5, 5 = 5, 5, 5, 5

				7 2 1		7 4 1
				4 5 6  =>	8 5 2
				9 8 3		9 6 3

				===============================================================================

				eg. i = 0, j = 0
				matrix[0][0], matrix[0][3], matrix[3][3], matrix[3][0] = matrix[3][0], matrix[0][0], matrix[0][3], matrix[3][3] => 1, 4, 16, 13 = 13, 1, 4, 16

				1  2  3  4		13 2  3  1
				5  6  7  8  =>	5  6  7  8
				9  10 11 12	    9  10 11 12
				13 14 15 16	    16 14 15 4

				eg. i = 0, j = 1
				matrix[0][1], matrix[1][3], matrix[3][2], matrix[3][0] = matrix[3][0], matrix[0][1], matrix[1][3], matrix[3][3] => 2, 8, 15, 9 = 9, 2, 8, 15

				13 2  3  1		13 9  3  1
				5  6  7  8  =>	5  6  7  2
				9  10 11 12	    15 10 11 12
				16 14 15 4	    16 14 8  4

				eg. i = 1, j = 0
				matrix[1][0], matrix[0][2], matrix[2][3], matrix[3][1] = matrix[3][1], matrix[1][0], matrix[0][2], matrix[2][3] => 5, 3, 12, 14 = 14, 5, 3, 12

				13 9  3  1		13 9  5  1
				5  6  7  2  =>	14 6  7  2
				15 10 11 12		15 10 11 3
				16 14 8  4	    16 12 8  4

				eg. i = 1, j = 1
				matrix[1][1], matrix[1][2], matrix[2][1], matrix[2][2] = matrix[2][2], matrix[1][1], matrix[1][2], matrix[2][1] => 6, 7, 11, 10 = 10, 6, 7, 11

				13 9  5  1		13 9  5  1
				14 6  7  2  =>	14 10 6  2
				15 10 11 3		15 11 7  3
				16 12 8  4	    16 12 8  4
			*/
			matrix[i][j], matrix[j][colNum-1-i], matrix[rowNum-1-i][colNum-1-j], matrix[rowNum-1-j][i] = matrix[rowNum-1-j][i], matrix[i][j], matrix[j][colNum-1-i], matrix[rowNum-1-i][colNum-1-j]
		}
	}

	return matrix
}

func Test_rotate1(t *testing.T) {
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
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			expected: expected{
				result: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			expected: expected{
				result: [][]int{
					{15, 13, 2, 5},
					{14, 3, 4, 1},
					{12, 6, 8, 9},
					{16, 7, 10, 11},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			rotate1(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_rotate2(t *testing.T) {
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
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			expected: expected{
				result: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			expected: expected{
				result: [][]int{
					{15, 13, 2, 5},
					{14, 3, 4, 1},
					{12, 6, 8, 9},
					{16, 7, 10, 11},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			rotate2(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_rotate3(t *testing.T) {
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
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			expected: expected{
				result: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			expected: expected{
				result: [][]int{
					{15, 13, 2, 5},
					{14, 3, 4, 1},
					{12, 6, 8, 9},
					{16, 7, 10, 11},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			rotate3(tc.args.matrix),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_rotate1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate1([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
	}
}

func Benchmark_rotate2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate2([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
	}
}

func Benchmark_rotate3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate3([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
	}
}
