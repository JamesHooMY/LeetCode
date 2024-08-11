package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/pacific-atlantic-water-flow/description/

// method 1 DFS
// 1) if the input is empty, return an empty slice
// 2) scan the left and right borders
// 3) scan the top and bottom borders
// 4) scan the matrix
// 5) scan the matrix again
// 6) return the result
// TC: O(M*N), SC: O(M*N), where M is the number of rows, N is the number of columns
// * this is the best solution for me currently
func pacificAtlantic1(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	// scan the left and right borders
	for r := 0; r < rows; r++ {
		pacificAtlanticDfs1(heights, r, 0, heights[r][0], pacific)
		pacificAtlanticDfs1(heights, r, cols-1, heights[r][cols-1], atlantic)
	}

	// scan the top and bottom borders
	for c := 0; c < cols; c++ {
		pacificAtlanticDfs1(heights, 0, c, heights[0][c], pacific)
		pacificAtlanticDfs1(heights, rows-1, c, heights[rows-1][c], atlantic)
	}

	result := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func pacificAtlanticDfs1(heights [][]int, r, c, prevHeight int, ocean [][]bool) {
	if r < 0 || r >= len(heights) || c < 0 || c >= len(heights[0]) || heights[r][c] < prevHeight || ocean[r][c] {
		return
	}

	ocean[r][c] = true // mark the visited node

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range directions {
		pacificAtlanticDfs1(heights, r+dir[0], c+dir[1], heights[r][c], ocean)
	}
}

// method 2 BFS
// 1) if the input is empty, return an empty slice
// 2) scan the left and right borders
// 3) scan the top and bottom borders
// 4) scan the matrix
// 5) scan the matrix again
// 6) return the result
// TC: O(M*N), SC: O(M*N), where M is the number of rows, N is the number of columns
func pacificAtlantic2(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(heights), len(heights[0])

	pacific := make([][]bool, rows)
	atlantic := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		pacific[i] = make([]bool, cols)
		atlantic[i] = make([]bool, cols)
	}

	// scan the left and right borders
	for r := 0; r < rows; r++ {
		pacificAtlanticBfs2(heights, r, 0, pacific)
		pacificAtlanticBfs2(heights, r, cols-1, atlantic)
	}

	// scan the top and bottom borders
	for c := 0; c < cols; c++ {
		pacificAtlanticBfs2(heights, 0, c, pacific)
		pacificAtlanticBfs2(heights, rows-1, c, atlantic)
	}

	result := [][]int{}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if pacific[r][c] && atlantic[r][c] {
				result = append(result, []int{r, c})
			}
		}
	}

	return result
}

func pacificAtlanticBfs2(heights [][]int, r, c int, ocean [][]bool) {
	rows, cols := len(heights), len(heights[0])

	queue := [][2]int{{r, c}} // * use array instead of slice [][]int{{r, c}} to improve performance
	ocean[r][c] = true // mark the visited node

	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			newR, newC := node[0]+dir[0], node[1]+dir[1]
			// * heights[node[0]][node[1]] is the height of the current node !!! Don't use heights[r][c] !!!
			if newR >= 0 && newR < rows && newC >= 0 && newC < cols && !ocean[newR][newC] && heights[newR][newC] >= heights[node[0]][node[1]] {
				queue = append(queue, [2]int{newR, newC}) // add the unvisited neighbor to the queue
				ocean[newR][newC] = true
			}
		}
	}
}

func Test_pacificAtlantic1(t *testing.T) {
	type args struct {
		heights [][]int
	}
	type expected struct {
		result [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 4},
					{1, 3},
					{1, 4},
					{2, 2},
					{3, 0},
					{3, 1},
					{4, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				heights: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 2},
					{1, 0},
					{1, 1},
					{1, 2},
					{2, 0},
					{2, 1},
					{2, 2},
				},
			},
		},
		{
			name: "3",
			args: args{
				heights: [][]int{
					{1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, pacificAtlantic1(tc.args.heights), tc.name)
	}
}

func Test_pacificAtlantic2(t *testing.T) {
	type args struct {
		heights [][]int
	}
	type expected struct {
		result [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				heights: [][]int{
					{1, 2, 2, 3, 5},
					{3, 2, 3, 4, 4},
					{2, 4, 5, 3, 1},
					{6, 7, 1, 4, 5},
					{5, 1, 1, 2, 4},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 4},
					{1, 3},
					{1, 4},
					{2, 2},
					{3, 0},
					{3, 1},
					{4, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				heights: [][]int{
					{1, 2, 3},
					{8, 9, 4},
					{7, 6, 5},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 2},
					{1, 0},
					{1, 1},
					{1, 2},
					{2, 0},
					{2, 1},
					{2, 2},
				},
			},
		},
		{
			name: "3",
			args: args{
				heights: [][]int{
					{1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, pacificAtlantic2(tc.args.heights), tc.name)
	}
}

// benchmark
func Benchmark_pacificAtlantic1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pacificAtlantic1([][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		})
	}
}

func Benchmark_pacificAtlantic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pacificAtlantic2([][]int{
			{1, 2, 2, 3, 5},
			{3, 2, 3, 4, 4},
			{2, 4, 5, 3, 1},
			{6, 7, 1, 4, 5},
			{5, 1, 1, 2, 4},
		})
	}
}
