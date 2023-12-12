package hard

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/palindrome-pairs/

// method 1 brute force
// 1) use nested loop to check if the concatenation of two words is palindrome
// 2) if yes, then add the index pair to result
// 3) finally, return the result
// TC = O(N^2 * K), SC = O(1), N is the length of words, K is the average length of word
func palindromePairs1(words []string) [][]int {
	result := [][]int{}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPalindrome(words[i] + words[j]) {
				result = append(result, []int{i, j})
			}
			if isPalindrome(words[j] + words[i]) {
				result = append(result, []int{j, i})
			}
		}
	}

	return result
}

// method 2 hash table
// 1) use a map to store the words (key: word, value: index)
// 2) use a slice to store the result
// 3) for each word, check if it is palindrome, if yes, then check if there is an empty string in the map, if yes, then add the index pair to result
// 4) for each word, check if the reverse of it is in the map, if yes, then add the index pair to result
// 5) for each word, split it into two parts, left and right, check if left is palindrome, if yes, then check if there is a reverse of right in the map, if yes, then add the index pair to result
// 6) for each word, split it into two parts, left and right, check if right is palindrome, if yes, then check if there is a reverse of left in the map, if yes, then add the index pair to result
// 7) finally, return the result
// TC = O(N * K^2), SC = O(N), N is the length of words, K is the average length of word
// * this is the best solution for me currently
func palindromePairs2(words []string) [][]int {
	result := [][]int{}

	wordIdxMap := make(map[string]int) // key: word, value: index
	for i := 0; i < len(words); i++ {
		wordIdxMap[words[i]] = i
	}

	for i := 0; i < len(words); i++ {
		word := words[i]
		wordLen := len(word)

		// check if the word is palindrome
		if isPalindrome(word) {
			// check if there is an empty string in the map
			if index, exist := wordIdxMap[""]; exist && index != i {
				result = append(result, []int{i, index}) // word + ""
				result = append(result, []int{index, i}) // "" + word

			}
		}

		// check if the reverse of the word is in the map
		reverseWord := reverseString(word)
		if index, exist := wordIdxMap[reverseWord]; exist && index != i {
			result = append(result, []int{i, index}) // word + reverse(word)
		}

		for j := 1; j < wordLen; j++ {
			leftWord := word[:j]
			rightWord := word[j:]

			// check if left is palindrome, if yes, then check if there is a reverse of right in the map
			if isPalindrome(leftWord) {
				reverseRight := reverseString(rightWord)
				if index, ok := wordIdxMap[reverseRight]; ok {
					if index != i {
						result = append(result, []int{index, i}) // reverseRight + word
					}
				}
			}

			// check if right is palindrome, if yes, then check if there is a reverse of left in the map
			if isPalindrome(rightWord) {
				reverseLeft := reverseString(leftWord)
				if index, ok := wordIdxMap[reverseLeft]; ok {
					if index != i {
						result = append(result, []int{i, index}) // word + reverseLeft
					}
				}
			}
		}
	}

	return result
}

func isPalindrome(s string) bool {
	sLen := len(s)
	if sLen == 0 {
		return false
	}

	left, right := 0, sLen-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func reverseString(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}

	// "abcde" -> "edcba"
	result := []byte(s)
	left, right := 0, sLen-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}

	return string(result)
}

func Test_palindromePairs1(t *testing.T) {
	type args struct {
		words []string
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
				words: []string{"abcd", "dcba", "lls", "s", "sssll"},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
			},
		},
		{
			name: "2",
			args: args{
				words: []string{"bat", "tab", "cat"},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}},
			},
		},
		{
			name: "3",
			args: args{
				words: []string{"a", ""},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			palindromePairs1(tc.args.words),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_palindromePairs2(t *testing.T) {
	type args struct {
		words []string
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
				words: []string{"abcd", "dcba", "lls", "s", "sssll"},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}},
			},
		},
		{
			name: "2",
			args: args{
				words: []string{"bat", "tab", "cat"},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}},
			},
		},
		{
			name: "3",
			args: args{
				words: []string{"a", ""},
			},
			expected: expected{
				result: [][]int{{0, 1}, {1, 0}},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			palindromePairs2(tc.args.words),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

var words = []string{"abcd", "dcba", "lls", "s", "sssll", "abcd", "dcba", "lls", "s", "sssll"}

// benchmark
func Benchmark_palindromePairs1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromePairs1(words)
	}
}

func Benchmark_palindromePairs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromePairs2(words)
	}
}
