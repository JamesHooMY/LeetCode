package util_test

import (
	"fmt"
	"testing"

	"leetcode/string/util"

	"github.com/stretchr/testify/assert"
)

func Test_IsAlphanumeric(t *testing.T) {
	type args struct {
		c rune
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
				c: rune('a'),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "2",
			args: args{
				c: rune('A'),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "3",
			args: args{
				c: rune('0'),
			},
			expected: expected{
				result: true,
			},
		},
		{
			name: "4",
			args: args{
				c: rune(' '),
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "5",
			args: args{
				c: rune('!'),
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
			util.IsAlphanumeric(tc.args.c),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_ToLowerCase(t *testing.T) {
	type args struct {
		c rune
	}
	type expected struct {
		result rune
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
				c: rune('a'),
			},
			expected: expected{
				result: rune('a'),
			},
		},
		{
			name: "2",
			args: args{
				c: rune('A'),
			},
			expected: expected{
				result: rune('a'),
			},
		},
		{
			name: "3",
			args: args{
				c: rune('0'),
			},
			expected: expected{
				result: rune('0'),
			},
		},
		{
			name: "4",
			args: args{
				c: rune(' '),
			},
			expected: expected{
				result: rune(' '),
			},
		},
		{
			name: "5",
			args: args{
				c: rune('!'),
			},
			expected: expected{
				result: rune('!'),
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			util.ToLowerCase(tc.args.c),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}