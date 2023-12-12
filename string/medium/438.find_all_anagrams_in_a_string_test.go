package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

// method 1 sliding window + hash table
// 1) use a pCharCountMap to store the characters of p (key: character, value: count)
// 2) use a wCharCountMap to store the characters of s (key: character, value: count), the size of the sliding windows is len(p), initial the wCharCountMap with the first len(p) characters of s
// 3) initial sliding window, the size of the window is len(p), use two pointers, left and right, left = 0, right = len(p) - 1
// 4) for each iteration, check if the wCharCountMap is an anagram of pCharCountMap, if yes, then append the left to the result
// 5) remove the character from the left of the window, move window to right, add the character of the right of the window to the map
// 6) finally, return the result
// TC = O(N), SC = O(N)
func findAnagrams1(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}

	// use a map to store the characters of p
	pCharCountMap := make(map[rune]int) // key: character, value: count
	for _, char := range p {
		pCharCountMap[char]++
	}

	// use a map to store the characters of s in the current window
	/*
		p = "abc"
		pCharCountMap := {
			'a': 1,
			'b': 1,
			'c': 1,
		}

		s = "cbaebabacd"
		wCharCountMap := {
			'c': 1,
			'b': 1,
			'a': 1,
		}
	*/
	wCharCountMap := make(map[rune]int) // key: character, value: count
	// the size of the sliding windows is len(p), initial the wCharCountMap with the first len(p) characters of s
	for i := 0; i < len(p); i++ {
		wCharCountMap[rune(s[i])]++
	}

	// initial sliding window, the size of the window is len(p)
	left := 0
	right := len(p) - 1

	for right < len(s) {
		if isAnagramMap(wCharCountMap, pCharCountMap) {
			result = append(result, left)
		}

		// remove the character from the left of the window
		wCharCountMap[rune(s[left])]--

		// move window to right
		left++
		right++

		// add the character of the right of the window to the map
		if right < len(s) {
			wCharCountMap[rune(s[right])]++
		}
	}

	return result
}

func isAnagramMap(wCharCountMap map[rune]int, pCharCountMap map[rune]int) bool {
	for char, count := range pCharCountMap {
		if wCharCountMap[char] != count {
			return false
		}
	}

	return true
}

// method 2 sliding window + array, performance and memory usage are better than method 1
// the logic thinking is same as method 1, but use array instead of map
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func findAnagrams2(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}

	// use an array to store the characters of p
	pArr := [26]int{} // index: character - 'a', value: count
	for _, char := range p {
		pArr[char-'a']++
	}

	// use an array to store the characters of s
	sArr := [26]int{} // index: character - 'a', value: count
	// the size of the sliding windows is len(p), initial the sArr with the first len(p) characters of s
	for i := 0; i < len(p); i++ {
		sArr[s[i]-'a']++
	}

	// initial sliding window, the size of the window is len(p)
	left := 0
	right := len(p) - 1

	for right < len(s) {
		if isAnagramArr(sArr, pArr) {
			result = append(result, left)
		}

		// remove the character from the left of the window
		sArr[s[left]-'a']--

		// move window to right
		left++
		right++

		// add the character of the right of the window to the map
		if right < len(s) {
			sArr[s[right]-'a']++
		}
	}

	return result
}

func isAnagramArr(sArr [26]int, pArr [26]int) bool {
	for i := 0; i < 26; i++ {
		if sArr[i] != pArr[i] {
			return false
		}
	}

	return true
}

func Test_findAnagrams1(t *testing.T) {
	type args struct {
		s string
		p string
	}
	type expected struct {
		result []int
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
				s: "cbaebabacd",
				p: "abc",
			},
			expected: expected{
				result: []int{0, 6},
			},
		},
		{
			name: "2",
			args: args{
				s: "abab",
				p: "ab",
			},
			expected: expected{
				result: []int{0, 1, 2},
			},
		},
		{
			name: "3",
			args: args{
				s: "aaaaaaaaaa",
				p: "aaaaaaaaaaaaa",
			},
			expected: expected{
				result: []int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findAnagrams1(tc.args.s, tc.args.p),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_findAnagrams2(t *testing.T) {
	type args struct {
		s string
		p string
	}
	type expected struct {
		result []int
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
				s: "cbaebabacd",
				p: "abc",
			},
			expected: expected{
				result: []int{0, 6},
			},
		},
		{
			name: "2",
			args: args{
				s: "abab",
				p: "ab",
			},
			expected: expected{
				result: []int{0, 1, 2},
			},
		},
		{
			name: "3",
			args: args{
				s: "aaaaaaaaaa",
				p: "aaaaaaaaaaaaa",
			},
			expected: expected{
				result: []int{},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			findAnagrams2(tc.args.s, tc.args.p),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
