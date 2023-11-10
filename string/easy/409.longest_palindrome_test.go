package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-palindrome/

// method 1 hash table
// 1) use a hash table to store the character and its count
// 2) use a variable to store the result
// 3) use hasOdd to store if there is an odd count character
// 4) for each character in the string, get its count from the hash table
// 5) result += count / 2 * 2, this is to get the even number no matter count is even or odd
// 6) if count % 2 != 0, then hasOdd = true
// 7) finally, if hasOdd is true, then result++
// 8) return result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func longestPalindrome1(s string) int {
	charMap := map[rune]int{} // key: character, value: count
	for _, v := range s {
		charMap[v]++
	}

	result := 0
	hasOdd := false
	for _, v := range charMap {
		// * this is the key point, get the even number no matter v is even or odd
		// for example, if v is 3, then v / 2 * 2 = 1 * 2 = 2, if v is 4, then v / 2 * 2 = 2 * 2 = 4
		result += v / 2 * 2

		if v%2 != 0 {
			hasOdd = true
		}
	}

	// the odd character can be placed in the middle of the palindrome
	if hasOdd {
		result++
	}

	return result
}

func Test_longestPalindrome1(t *testing.T) {
	type args struct {
		s string
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
				s: "abccccdd",
			},
			expected: expected{
				result: 7,
			},
		},
		{
			name: "2",
			args: args{
				s: "a",
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				s: "bb",
			},
			expected: expected{
				result: 2,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestPalindrome1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
