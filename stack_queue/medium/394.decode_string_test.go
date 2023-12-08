package medium

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/decode-string/description/

// method 1 two stacks
// 1) use two stacks to store the previous string and the multiplier of current string
// 2) if the current character is number, then calculate the multiplier
// 3) if the current character is '[', then push the current string and multiplier into stacks
// 4) if the current character is ']', then pop the previous strin
// 5) if the current character is letter, then store the current string
// 6) finally, the current string is the result
// TC = O(N), SC = O(N)
func decodeString1(s string) string {
	strStack := []string{} // store the previous string
	multStack := []int{}   // store the multiple of current string

	currentStr := ""
	mult := 0 // multiplier of current string

	for _, v := range s {
		if v >= '0' && v <= '9' {
			mult = mult*10 + int(v-'0')
		} else if v == '[' {
			strStack = append(strStack, currentStr) // store the current string
			currentStr = ""                         // reset the current string

			multStack = append(multStack, mult) // store the multiplier of current string
			mult = 0                            // reset the multiplier
		} else if v == ']' {
			prevStr := strStack[len(strStack)-1]  // get the previous string
			strStack = strStack[:len(strStack)-1] // pop the previous string

			curMult := multStack[len(multStack)-1]   // get the multiplier of current string
			multStack = multStack[:len(multStack)-1] // pop the multiplier of current string

			currentStr = prevStr + strings.Repeat(currentStr, curMult)
		} else {
			currentStr += string(v) // store the current string
		}
	}

	return currentStr
}

// method 2 recursion
// 1) use a pointer to store the current index of string
// 2) if the current character is number, then calculate the multiplier
// 3) if the current character is '[', then call the recursion function to get the current string
// 4) if the current character is ']', then return the current string
// 5) if the current character is letter, then store the current string
// 6) finally, the current string is the result
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func decodeString2(s string) string {
	i := 0 // pointer to store the current index of string

	return decodeString2Helper(s, &i)
}

// * key point: use a pointer to store the current index of string
func decodeString2Helper(s string, i *int) string {
	currentStr := ""
	mult := 0 // multiplier of current string

	for *i < len(s) {
		v := s[*i]
		*i++

		if v >= '0' && v <= '9' {
			mult = mult*10 + int(v-'0')
		} else if v == '[' {
			str := decodeString2Helper(s, i)
			currentStr += strings.Repeat(str, mult)

			mult = 0 // reset the multiplier
		} else if v == ']' {
			return currentStr
		} else {
			currentStr += string(v) // store the current string
		}
	}

	return currentStr
}

func Test_decodeString1(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result string
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
				s: "3[a]2[bc]",
			},
			expected: expected{
				result: "aaabcbc",
			},
		},
		{
			name: "2",
			args: args{
				s: "3[a2[c]]",
			},
			expected: expected{
				result: "accaccacc",
			},
		},
		{
			name: "3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			expected: expected{
				result: "abcabccdcdcdef",
			},
		},
		{
			name: "4",
			args: args{
				s: "abc3[cd]xyz",
			},
			expected: expected{
				result: "abccdcdcdxyz",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			decodeString1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_decodeString2(t *testing.T) {
	type args struct {
		s string
	}
	type expected struct {
		result string
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
				s: "3[a]2[bc]",
			},
			expected: expected{
				result: "aaabcbc",
			},
		},
		{
			name: "2",
			args: args{
				s: "3[a2[c]]",
			},
			expected: expected{
				result: "accaccacc",
			},
		},
		{
			name: "3",
			args: args{
				s: "2[abc]3[cd]ef",
			},
			expected: expected{
				result: "abcabccdcdcdef",
			},
		},
		{
			name: "4",
			args: args{
				s: "abc3[cd]xyz",
			},
			expected: expected{
				result: "abccdcdcdxyz",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			decodeString2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_decodeString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeString1("3[a]2[bc]")
	}
}

func Benchmark_decodeString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeString2("3[a]2[bc]")
	}
}