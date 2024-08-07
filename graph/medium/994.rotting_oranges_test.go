package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/rotting-oranges/description/

// method 1 BFS
// use a queue to store rotten oranges
// count fresh oranges
// if no fresh oranges, return 0
// directions: up, down, left, right
// for each rotten orange, check its neighbors
// if the neighbor is fresh, change it to rotten, add it to queue, and decrease fresh count
// return minutes
// TC: O(m*n), SC: O(m*n), where m is the number of rows, n is the number of columns
// * this is the best solution for me currently
func orangesRotting1(grid [][]int) int {
	minutes := 0
	freshCount := 0

	rows, cols := len(grid), len(grid[0])

	// count fresh oranges, and add rotten oranges to queue
	queue := [][]int{} // store rotten oranges
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				freshCount++
			} else if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}
		}
	}

	// if no fresh oranges, return 0
	if freshCount == 0 {
		return 0
	}

	// directions
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 && freshCount > 0 { // * freshCount > 0 is important !!!
		minutes++
		size := len(queue) // * this is important, because the length of queue will change

		for i := 0; i < size; i++ {
			rotten := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				newR, newC := rotten[0]+dir[0], rotten[1]+dir[1]
				if newR >= 0 && newR < rows && newC >= 0 && newC < cols && grid[newR][newC] == 1 {
					grid[newR][newC] = 2
					freshCount--
					queue = append(queue, []int{newR, newC})
				}
			}
		}
	}

	if freshCount > 0 {
		return -1
	}

	return minutes
}

func Test_orangesRotting1(t *testing.T) {
	type args struct {
		grid [][]int
	}
	type expected struct {
		result int
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
				grid: [][]int{
					{2, 1, 1},
					{1, 1, 0},
					{0, 1, 1},
				},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				grid: [][]int{
					{2, 1, 1},
					{0, 1, 1},
					{1, 0, 1},
				},
			},
			expected: expected{
				result: -1,
			},
		},
		{
			name: "3",
			args: args{
				grid: [][]int{
					{0, 2},
				},
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "4",
			args: args{
				grid: [][]int{
					{2, 2, 2, 1, 1},
					{2, 2, 1, 2, 1},
					{2, 2, 1, 1, 1},
				},
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, orangesRotting1(tc.args.grid))
	}
}
