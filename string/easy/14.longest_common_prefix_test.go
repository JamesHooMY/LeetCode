package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-common-prefix/description/

// method 1
// 1) use the first string as the prefix, prefix = strs[0]
// 2) use for loop to scan the strs,
// 3) meanwhile, use another while loop with a index j to scan the prefix and strs[i] from the beginning
// 4) if prefix[j] != strs[i][j], then remove the last character of prefix, prefix = prefix[:j]
// 5) if prefix == "", then return ""
// 6) finally, return prefix
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// use the first string as the prefix
	prefix := strs[0]

	// use for loop to scan the strs
	// if the prefix is not the prefix of strs[i], then remove the last character of prefix
	// if prefix == "", then return ""
	for i := 1; i < len(strs); i++ {
		// j is the index to scan the prefix and strs[i] from the beginning
		// * prefix must appear in all elements of strs with the same order, length, and position
		j := 0
		for j < len(prefix) && j < len(strs[i]) && prefix[j] == strs[i][j] {
			j++
		}

		// * the description of this problem indicates that the prefix must appear in all elements of strs
		// so if one element of strs does not contain the prefix, then return ""
		if j == 0 {
			return ""
		}

		// only keep the common prefix during the iteration
		prefix = prefix[:j]
	}

	return prefix
}

func Test_longestCommonPrefix1(t *testing.T) {
	type args struct {
		strs []string
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
				strs: []string{"flower", "flow", "flight"},
			},
			expected: expected{
				result: "fl",
			},
		},
		{
			name: "2",
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			expected: expected{
				result: "",
			},
		},
		{
			name: "3",
			args: args{
				strs: []string{"a"},
			},
			expected: expected{
				result: "a",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestCommonPrefix1(tc.args.strs),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
