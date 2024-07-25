package medium

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"leetcode/util"
)
// https://leetcode.com/problems/01-matrix/description/

// method 1 iterative BFS (top-down)
// 1) use dist to store the distance from 0 to 1, and set all dist to -1 for initialization
// 2) push all 0s into queue, and set dist to 0, because 0 is the starting point
// 3) iterate queue while queue is not empty
// 4) pop cell from queue, iterate directions, calculate newX and newY
// 5) if newX and newY is valid, and dist[newX][newY] is -1, set dist[newX][newY] = dist[x][y] + 1, and push newX and newY into queue
// 6) return dist
// TC = O(M*N), SC = O(M*N), M is the number of rows, N is the number of cols
func updateMatrix1(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows == 0 || cols == 0 {
		return mat
	}

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = -1 // * set all dist to -1 for initialization
		}
	}

	queue := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
				dist[i][j] = 0 // * set dist to 0 for 0s
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:] // pop
		x, y := cell[0], cell[1]

		// iterate directions around cell
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && dist[newX][newY] == -1 {
				dist[newX][newY] = dist[x][y] + 1        // * this is the key point for distance accumulation
				queue = append(queue, []int{newX, newY}) // * push newX and newY into queue
			}
		}
	}

	return dist
}

// method 2 iterative BFS (top-down) combine the initialization of dist and pushing 0s into queue with inplace
// 1) push all 0s into queue, and set other cells to -1
// 2) iterate queue while queue is not empty
// 3) pop cell from queue, iterate directions, calculate newX and newY
// 4) if newX and newY is valid, and mat[newX][newY] is -1, set mat[newX][newY] = mat[x][y] + 1, and push newX and newY into queue
// 5) return mat
// TC = O(M*N), SC = O(M*N), M is the number of rows, N is the number of cols
func updateMatrix2(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows == 0 || cols == 0 {
		return mat
	}

	queue := [][]int{} // * push all 0s into queue
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1 // * set all 1s to -1 for initialization
			}
		}
	}

	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // up, down, left, right
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:] // pop
		x, y := cell[0], cell[1]

		// iterate directions around cell
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols && mat[newX][newY] == -1 {
				mat[newX][newY] = mat[x][y] + 1          // * this is the key point for distance accumulation
				queue = append(queue, []int{newX, newY}) // * push newX and newY into queue
			}
		}
	}

	return mat
}

// method 3 DP with Two Passes (Top-Left to Bottom-Right, Bottom-Right to Top-Left)
// 1) initialize dist with max value for all cells, because we need to find the minimum distance, and avoid overflow
// 2) iterate from top-left to bottom-right, compare with up cell and left cell
// 3) iterate from bottom-right to top-left, compare with down cell and right cell
// 4) return dist
// TC = O(M*N), SC = O(M*N), M is the number of rows, N is the number of cols
func updateMatrix3(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows == 0 || cols == 0 {
		return mat
	}

	// * initialize dp with max value
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32 // * max value, avoid overflow !!!
		}
	}

	// * iterate from top-left -> bottom-right
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					dist[i][j] = util.Min(dist[i][j], dist[i-1][j]+1) // * compare with up cell
				}

				if j > 0 {
					dist[i][j] = util.Min(dist[i][j], dist[i][j-1]+1) // * compare with left cell
				}
			}
		}
	}

	// * iterate from bottom-right -> top-left
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if i < rows-1 {
				dist[i][j] = util.Min(dist[i][j], dist[i+1][j]+1) // * compare with down cell
			}

			if j < cols-1 {
				dist[i][j] = util.Min(dist[i][j], dist[i][j+1]+1) // * compare with right cell
			}
		}
	}

	return dist
}

// method 4 DP with Two Passes optimized in-place
// 1) iterate from top-left to bottom-right, compare with up cell and left cell
// 2) iterate from bottom-right to top-left, compare with down cell and right cell
// 3) return mat
// TC = O(M*N), SC = O(1), M is the number of rows, N is the number of cols
// * this is the best solution for me currently
func updateMatrix4(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	if rows == 0 || cols == 0 {
		return mat
	}

	// maxDist := rows + cols // * max distance, maybe overflow
	maxDist := math.MaxInt32 // * max distance, avoid overflow !!!

	// * iterate from top-left -> bottom-right
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] != 0 {
				top, left := maxDist, maxDist

				if i > 0 {
					top = mat[i-1][j]
				}

				if j > 0 {
					left = mat[i][j-1]
				}

				mat[i][j] = util.Min(top, left) + 1
			}
		}
	}

	// * iterate from bottom-right -> top-left
	for i := rows - 1; i >= 0; i-- {
		for j := cols - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				down, right := maxDist, maxDist

				if i < rows-1 {
					down = mat[i+1][j]
				}

				if j < cols-1 {
					right = mat[i][j+1]
				}

				mat[i][j] = util.Min(mat[i][j], util.Min(down, right)+1)
			}
		}
	}

	return mat
}

func Test_updateMatrix1(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
				},
			},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
					{2, 3, 2},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			updateMatrix1(tc.args.mat),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_updateMatrix2(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
				},
			},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
					{2, 3, 2},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			updateMatrix2(tc.args.mat),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_updateMatrix3(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
				},
			},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
					{2, 3, 2},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			updateMatrix3(tc.args.mat),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_updateMatrix4(t *testing.T) {
	type args struct {
		mat [][]int
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
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
		},
		{
			name: "2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
				},
			},
		},
		{
			name: "3",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			expected: expected{
				result: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 2, 1},
					{2, 3, 2},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			updateMatrix4(tc.args.mat),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_updateMatrix1(b *testing.B) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		updateMatrix1(mat)
	}
}

func Benchmark_updateMatrix2(b *testing.B) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		updateMatrix2(mat)
	}
}

func Benchmark_updateMatrix3(b *testing.B) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		updateMatrix3(mat)
	}
}

func Benchmark_updateMatrix4(b *testing.B) {
	mat := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	for i := 0; i < b.N; i++ {
		updateMatrix4(mat)
	}
}
