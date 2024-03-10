package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-islands/description/

// method 1: DFS
// 1) use a variable to store the number of islands
// 2) use a 2D array to store whether a cell has been visited
// 3) traverse the grid
// 4) if the cell is an island and has not been visited, traverse it and increment the number of islands
// TC = O(M*N), SC = O(M*N), where M is the number of rows and N is the number of columns
// * this is the best solution for me currently
func numIslands1(grid [][]byte) int {
	rowNum := len(grid)
	colNum := len(grid[0])

	// use a variable to store the number of islands
	islandNum := 0

	// use a 2D array to store whether a cell has been visited
	visited := make([][]bool, rowNum)
	for i := 0; i < rowNum; i++ {
		/*
			example:
				rowNum = 2
				colNum = 2
				visited = [
					[false, false],
					[false, false],
				]
		*/
		visited[i] = make([]bool, colNum)
	}

	// traverse the grid
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			// if the cell is an island and has not been visited, traverse it and increment the number of islands
			if grid[i][j] == '1' && !visited[i][j] {
				traverse_DFS(grid, visited, i, j, rowNum, colNum)
				islandNum++
			}
		}
	}

	return islandNum
}

// helper function to check whether a cell is valid, not over the boundary
func isValidCell(rowIdx, colIdx, rowNum, colNum int) bool {
	return rowIdx >= 0 && rowIdx < rowNum && colIdx >= 0 && colIdx < colNum
}

// helper function to traverse the grid
func traverse_DFS(grid [][]byte, visited [][]bool, rowIdx, colIdx, rowNum, colNum int) {
	// if the cell is invalid or has been visited or is not an island, return
	if !isValidCell(rowIdx, colIdx, rowNum, colNum) || visited[rowIdx][colIdx] || grid[rowIdx][colIdx] == '0' {
		return
	}

	// mark the cell as visited
	visited[rowIdx][colIdx] = true

	// traverse the cell's four neighbors, i.e. top, bottom, left, right
	topIdx, btmIdx, leftIdx, rightIdx := rowIdx-1, rowIdx+1, colIdx-1, colIdx+1
	traverse_DFS(grid, visited, topIdx, colIdx, rowNum, colNum)   // top
	traverse_DFS(grid, visited, btmIdx, colIdx, rowNum, colNum)   // bottom
	traverse_DFS(grid, visited, rowIdx, leftIdx, rowNum, colNum)  // left
	traverse_DFS(grid, visited, rowIdx, rightIdx, rowNum, colNum) // right
}

// method 2: BFS
// 1) use a variable to store the number of islands
// 2) use a 2D array to store whether a cell has been visited
// 3) traverse the grid
// 4) if the cell is an island and has not been visited, traverse it and increment the number of islands
// TC = O(M*N), SC = O(M*N), where M is the number of rows and N is the number of columns
func numIslands2(grid [][]byte) int {
	rowNum := len(grid)
	colNum := len(grid[0])

	// use a variable to store the number of islands
	islandNum := 0

	// use a 2D array to store whether a cell has been visited
	visited := make([][]bool, rowNum)
	for i := 0; i < rowNum; i++ {
		visited[i] = make([]bool, colNum)
	}

	// traverse the grid
	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			// if the cell is an island and has not been visited, traverse it and increment the number of islands
			if grid[i][j] == '1' && !visited[i][j] {
				traverse_BFS(grid, visited, i, j, rowNum, colNum)
				islandNum++
			}
		}
	}

	return islandNum
}

func traverse_BFS(grid [][]byte, visited [][]bool, rowIdx, colIdx, rowNum, colNum int) {
	// use a queue to store the cell's neighbors for BFS iteration
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{rowIdx, colIdx})

	// traverse the grid
	for len(queue) > 0 {
		// pop the first cell from the queue
		cell := queue[0]
		queue = queue[1:]

		// get the cell's row and column
		row := cell[0]
		col := cell[1]

		// traverse the cell's four neighbors, i.e. top, bottom, left, right
		topIdx, btmIdx, leftIdx, rightIdx := row-1, row+1, col-1, col+1
		if isValidCell(topIdx, colIdx, rowNum, colNum) && !visited[topIdx][col] && grid[topIdx][col] == '1' {
			visited[topIdx][col] = true
			queue = append(queue, [2]int{topIdx, col})
		}
		if isValidCell(btmIdx, colIdx, rowNum, colNum) && !visited[btmIdx][col] && grid[btmIdx][col] == '1' {
			visited[btmIdx][col] = true
			queue = append(queue, [2]int{btmIdx, col})
		}
		if isValidCell(rowIdx, leftIdx, rowNum, colNum) && !visited[row][leftIdx] && grid[row][leftIdx] == '1' {
			visited[row][leftIdx] = true
			queue = append(queue, [2]int{rowIdx, leftIdx})
		}
		if isValidCell(rowIdx, rightIdx, rowNum, colNum) && !visited[row][rightIdx] && grid[row][rightIdx] == '1' {
			visited[row][rightIdx] = true
			queue = append(queue, [2]int{rowIdx, rightIdx})
		}
	}
}

func Test_numIslands1(t *testing.T) {
	type args struct {
		grid [][]byte
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
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			numIslands1(tc.args.grid),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_numIslands2(t *testing.T) {
	type args struct {
		grid [][]byte
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
				grid: [][]byte{
					{'1', '1', '1', '1', '0'},
					{'1', '1', '0', '1', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '0', '0', '0'},
				},
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "2",
			args: args{
				grid: [][]byte{
					{'1', '1', '0', '0', '0'},
					{'1', '1', '0', '0', '0'},
					{'0', '0', '1', '0', '0'},
					{'0', '0', '0', '1', '1'},
				},
			},
			expected: expected{
				result: 3,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			numIslands2(tc.args.grid),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_numIslands1(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	for i := 0; i < b.N; i++ {
		numIslands1(grid)
	}
}

func Benchmark_numIslands2(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	for i := 0; i < b.N; i++ {
		numIslands2(grid)
	}
}
