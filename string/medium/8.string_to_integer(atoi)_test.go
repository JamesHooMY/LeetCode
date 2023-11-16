package medium

import (
	"fmt"
	"math"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/string-to-integer-atoi/description/

// method 1
// 1) trim the string
// 2) check the sign
// 3) use a for loop to iterate the string, if the character is not a digit, then break
// 4) use a variable to store the result, check the overflow for each iteration, then update the result
// 5) return the result * sign
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func myAtoi1(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// use a variable to store the index
	index := 0
	// use a variable to store the sign
	sign := 1
	// check the sign
	if s[index] == '+' || s[index] == '-' {
		if s[index] == '-' {
			sign = -1
		}
		index++
	}

	// use a variable to store the result
	result := 0
	// use a for loop to iterate the string
	for i := index; i < len(s); i++ {
		// if the character is not a digit, then break
		if s[i] < '0' || s[i] > '9' {
			break
		}

		// * this is the key point, check the overflow
		/*
			math.MaxInt32 = 2147483647
			math.MinInt32 = -2147483648

			if result > math.MaxInt32 / 10, then result * 10 will overflow
			example: result = 214748364, then result * 10 = 2147483640, which is overflow

			if result == math.MaxInt32 / 10 && curNum > math.MaxInt32 % 10, then result * 10 + curNum will overflow
			example: result = 214748364, curNum = 8, then result * 10 + curNum = 2147483648, which is overflow
		*/
		curNum := int(s[i] - '0')
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && curNum > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}

		// * this is the key point, update the result
		result = result*10 + curNum
	}

	return result * sign
}

func Test_myAtoi1(t *testing.T) {
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
				s: "42",
			},
			expected: expected{
				result: 42,
			},
		},
		{
			name: "2",
			args: args{
				s: "   -42",
			},
			expected: expected{
				result: -42,
			},
		},
		{
			name: "3",
			args: args{
				s: "4193 with words",
			},
			expected: expected{
				result: 4193,
			},
		},
		{
			name: "4",
			args: args{
				s: "words and 987",
			},
			expected: expected{
				result: 0,
			},
		},
		{
			name: "5",
			args: args{
				s: "-91283472332",
			},
			expected: expected{
				result: -2147483648,
			},
		},
		{
			name: "6",
			args: args{
				s: "2147483648",
			},
			expected: expected{
				result: 2147483647,
			},
		},
		{
			name: "7",
			args: args{
				s: "9223372036854775808",
			},
			expected: expected{
				result: 2147483647,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			myAtoi1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
