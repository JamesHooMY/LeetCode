package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/word-search/description/

// method 1 DFS
// 1) use a for loop to scan the rows of the board
// 2) use a for loop to scan the columns of the board
// 3) call the dfs function to check if the word exists
// 4) dfs function is used to check if the word exists in the board
// 5) check the boundary and the character, if not match, return false
// 6) get the last character of the word, if match, return true
// 7) mark the character as visited
// 8) call the dfs function recursively to check the neighbors
// 9) reset the character
// 10) return the result
// 11) return the result of the dfs function
// TC: O(N*3^L), SC: O(L), where N is the number of cells in the board, L is the length of the word, each cell has 3 directions to go
func exist1(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if existDfs1(board, r, c, 0, word) {
				return true
			}
		}
	}

	return false
}

func existDfs1(board [][]byte, r, c, idx int, word string) bool {
	// check the boundary and the character
	if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) || board[r][c] != word[idx] {
		return false
	}

	// get the last character of the word
	if idx == len(word) - 1 {
		return true
	}

	temp := board[r][c]
	board[r][c] = '#' // * mark as visited

	result := existDfs1(board, r-1, c, idx+1, word) ||
		existDfs1(board, r+1, c, idx+1, word) ||
		existDfs1(board, r, c-1, idx+1, word) ||
		existDfs1(board, r, c+1, idx+1, word)

	board[r][c] = temp // * reset the character

	return result
}

func Test_exist1(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
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
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "SEE",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCB",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, exist1(tc.args.board, tc.args.word), tc.name)
	}
}
