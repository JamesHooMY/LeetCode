package easy

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"leetcode/string/util"

	"github.com/stretchr/testify/assert"
)

// method 1 two pointers
// 1) use two pointers, left and right
// 2) use while loop to scan the string, if left < right, then compare s[left] and s[right]
// 3) check if s[left] and s[right] are alphanumeric, if not, then left++ or right-- and continue
// 4) check if s[left] == s[right], if not, then turn them to lowercase and compare again
// 5) if s[left] != s[right], then return false, otherwise, left++, right--
// 6) finally, return true
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func isPalindrome1(s string) bool {
	// use two pointers, left and right
	left := 0
	right := len(s) - 1

	// use while loop to scan the string
	// if left < right, then compare s[left] and s[right]
	// if s[left] != s[right], then return false
	// otherwise, left++, right--
	for left < right {
		if !util.IsAlphanumeric(s[left]) {
			left++
			continue
		}

		if !util.IsAlphanumeric(s[right]) {
			right--
			continue
		}

		// * this is the key point, optimize the efficiency of the code by reducing the number of calls to ToLowerCase
		if s[left] == s[right] {
			left++
			right--
			continue
		}

		if !(util.ToLowerCase(s[left]) == util.ToLowerCase(s[right])) {
			return false
		}

		left++
		right--
	}

	return true
}

// method 2 two pointers, regular expression, this method take much more time than method 1
// 1) use regular expression to remove all non-alphanumeric characters
// 2) turn the string to lowercase
// 3) use two pointers, left and right
// 4) use while loop to scan the string, if left < right, then compare s[left] and s[right]
// 5) if s[left] != s[right], then return false, otherwise, left++, right--
// 6) finally, return true
// TC = O(N), SC = O(1)
func isPalindrome2(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(s)

	// use two pointers, left and right
	left := 0
	right := len(s) - 1

	// use while loop to scan the string
	// if left < right, then compare s[left] and s[right]
	// if s[left] != s[right], then return false
	// otherwise, left++, right--
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func Test_isPalindrome1(t *testing.T) {
	type args struct {
		s string
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
				s: "A man, a plan, a canal: Panama",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "race a car",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				s: " ",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "4",
			args: args{
				s: "0P",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isPalindrome1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_isPalindrome2(t *testing.T) {
	type args struct {
		s string
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
				s: "A man, a plan, a canal: Panama",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				s: "race a car",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				s: " ",
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "4",
			args: args{
				s: "0P",
			},
			expected: expected{
				result: false,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			isPalindrome2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_isPalindrome1(b *testing.B) {
	s := "A man, a plan, a canal: Panama"

	for i := 0; i < b.N; i++ {
		isPalindrome1(s)
	}
}

func Benchmark_isPalindrome2(b *testing.B) {
	s := "A man, a plan, a canal: Panama"

	for i := 0; i < b.N; i++ {
		isPalindrome2(s)
	}
}
