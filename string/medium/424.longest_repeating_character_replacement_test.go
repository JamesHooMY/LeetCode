package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-repeating-character-replacement/

// method 1 sliding window + hash table
// 1) use a charMap to store the characters of s (key: character, value: count) in the current window
// 2) use left and right pointer to represent the current window
// 3) use maxCount to store the max count of the same character in the current window
// 4) use maxLength to store the max length of the current window
// 5) use right pointer to move the window, and update the charMap and maxCount
// 6) if the current window size - maxCount > k, it means the other characters count should be replaced in current window is more than the k requirement, so we need to move the left pointer to shrink the window, and update the charMap, maxCount and maxLength
// 7) repeat step 5 and 6 until right pointer reach the end of s, and return maxLength
// TC = O(n), SC = O(n)
func characterReplacement1(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	// use a map to store the characters of s
	/*
		s = "AABABBA"
		sMap := {
			'A': 3,
			'B': 4,
		}
	*/
	sMap := make(map[byte]int) // key: character, value: count

	// initial sliding window
	left := 0
	right := 0
	maxCount := 0  // maxCount of character in current window
	maxLength := 0 // maxLength of current window

	for right < len(s) {
		sMap[s[right]]++
		maxCount = max(maxCount, sMap[s[right]])

		// right - left + 1 is the current window size, current window size - maxCount = the other characters count should be replaced in current window
		/*
				s: A A B A B B A
			index: 0 1 2 3 4 5 6
				k: 1

			right = 0, left = 0, sMap = {'A': 1}, maxCount = 1, result = 1
			right = 1, left = 0, sMap = {'A': 2}, maxCount = 2, result = 2
			right = 2, left = 0, sMap = {'A': 2, 'B': 1}, maxCount = 2, result = 3
			right = 3, left = 0, sMap = {'A': 3, 'B': 1}, maxCount = 3, result = 4
			right = 4, left = 0 => 1, sMap = {'A': 3, 'B': 2} => {'A': 2, 'B': 2}, maxCount = 3, result = 4
			* right = 4, this step remove the first 'A' in the current window (index 0 -> 4), and move the left pointer to the next position, shrink the window (index from 0 -> 4 shrink to 1 -> 4)

			right = 5, left = 1 => 2, sMap = {'A': 2, 'B': 3} => {'A': 1, 'B': 3}, maxCount = 3, result = 4
			* right = 5, this step remove the first 'A' in the current window (index 1 -> 5), and move the left pointer to the next position, shrink the window (index from 1 -> 5 shrink to 2 -> 5)

			right = 6, left = 2 => 3, sMap = {'A': 2, 'B': 3} => {'A': 2, 'B': 2}, maxCount = 3, result = 4
			* right = 6, this step remove the first 'B' in the current window (index 2 -> 6), and move the left pointer to the next position, shrink the window (index from 2 -> 6 shrink to 3 -> 6)
		*/
		if (right-left+1)-maxCount > k {
			sMap[s[left]]--
			left++
		}
		maxLength = max(maxLength, right-left+1)
		right++
	}

	return maxLength
}

// method 2 sliding window + array
// logic thinking is same as method 1, but use array instead of map
// TC = O(n), SC = O(n)
// * this is the best solution for me currently
func characterReplacement2(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	// use a array to store the characters of s
	charArr := [26]int{} // index: character - 'A', value: count

	// initial sliding window
	left := 0
	maxCount := 0  // maxCount of character in current window
	maxLength := 0 // maxLength of current window

	for right := 0; right < len(s); right++ {
		charArr[s[right]-'A']++
		maxCount = max(maxCount, charArr[s[right]-'A'])

		// right - left + 1 is the current window size, current window size - maxCount = the other characters count should be replaced in current window
		if (right-left+1)-maxCount > k {
			charArr[s[left]-'A']--
			left++
		}

		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func Test_characterReplacement1(t *testing.T) {
	type args struct {
		s string
		k int
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
				s: "ABAB",
				k: 2,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				s: "AABABBA",
				k: 0,
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
			characterReplacement1(tc.args.s, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_characterReplacement2(t *testing.T) {
	type args struct {
		s string
		k int
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
				s: "ABAB",
				k: 2,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				s: "AABABBA",
				k: 1,
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "3",
			args: args{
				s: "AABABBA",
				k: 0,
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
			characterReplacement2(tc.args.s, tc.args.k),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}