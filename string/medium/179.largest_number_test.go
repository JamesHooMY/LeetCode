package easy

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-number/description/

// method 1
// 1) convert nums to string
// 2) sort the string array, compare a + b and b + a (if a + b > b + a, then a should be in front of b, otherwise b should be in front of a)
// 3) if the first element is "0", then return "0"
// 4) join the string array
// TC = O(NlogN), SC = O(N)
// * this is the best solution for me currently
func largestNumber1(nums []int) string {
	// convert nums to string
	strs := make([]string, len(nums))
	for i, v := range nums {
		// strs[i] = fmt.Sprintf("%d", v)
		// * strconv.Itoa(v) is faster than fmt.Sprintf("%d", v)
		strs[i] = strconv.Itoa(v)
	}

	// sort the string array, TC = O(NlogN), SC = O(logN)
	// if a + b > b + a, then a should be in front of b, otherwise b should be in front of a
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// if the first element is "0", then return "0"
	if strs[0] == "0" {
		return "0"
	}

	// join the string array
	return strings.Join(strs, "")
}

func Test_largestNumber1(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{10, 2},
			},
			expected: expected{
				result: "210",
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			expected: expected{
				result: "9534330",
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{0},
			},
			expected: expected{
				result: "0",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			largestNumber1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_largestNumber1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		largestNumber1([]int{10, 2})
	}
}
