package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// method 1 sliding window + hash table
// 1) use two pointers, start and end
// 2) use a map to store the character and its index
// 3) use a variable to store the maxLength (end - start + 1)
// 4) for each character in the string, if the character is already in the map, then move start to the index of the character + 1
// 5) update the index of the character in the map, update end to the current index
// 6) finally, return the maxLength
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	// use two pointers, left and right
	start := 0
	end := 0

	// use a map to store the character and its index
	// * benchmark shows that the performance of using rune is better than using byte
	charMap := map[rune]int{} // key: character, value: index
	// charMap := map[byte]int{} // key: character, value: index

	// use a variable to store the maxLength
	maxLength := 0

	for i := 0; i < len(s); i++ {
		// if index, exist := charMap[s[i]]; exist {
		if index, exist := charMap[rune(s[i])]; exist {
			// index >= start instead of index > start, because of the case "bbbbb"
			if index >= start {
				start = index + 1
			}
		}

		// charMap[s[i]] = i
		charMap[rune(s[i])] = i
		end = i

		if maxLength < end-start+1 {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

func Test_lengthOfLongestSubstring1(t *testing.T) {
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
				s: "abcabcbb",
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				s: "bbbbb",
			},
			expected: expected{
				result: 1,
			},
		},
		{
			name: "3",
			args: args{
				s: "pwwkew",
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
			lengthOfLongestSubstring1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_lengthOfLongestSubstring1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring1("abcabcbb")
	}
}
