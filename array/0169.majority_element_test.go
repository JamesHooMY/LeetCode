package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// method 1
// use two for loop, one map
// 1) first for loop, use map to store each num and its count
// 2) second for loop, check each num in map, if count > len(nums)/2, return num
// TC = O(2N), SC = O(N)
func majorityElement1(nums []int) int {
	numCountMap := map[int]int{}
	for _, num := range nums {
		numCountMap[num]++
	}

	for num, count := range numCountMap {
		// count must be greater than len(nums)/2, then return num
		if count > len(nums)/2 {
			return num
		}
	}

	return -1
}

// method 2 Boyer–Moore majority vote algorithm
// 1) use two variable, count and candidate
// 2) first for loop, if count == 0, set candidate = num
// 3) if num == candidate, count++
// TC = O(N), SC = O(1)
// * this method is restricted to slices containing only two types of num, one of which is the majority
func majorityElement2(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	if count == 0 {
		return -1
	}

	return candidate
}

// method 2 Boyer–Moore majority vote algorithm, with check
// 1) use two variable, count and candidate
// 2) first for loop, if count == 0, set candidate = num
// 3) if num == candidate, count++
// TC = O(2N), SC = O(1)
// * this is the best solution to fit all situation and optimize the SC to 0(1) for me by now
func majorityElement3(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	if count > len(nums)/2 {
		return candidate
	}

	return -1
}

func Test_majorityElement1(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 2, 3},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 3, 4, 3, 4, 1, 4, 3, 3, 3, 3, 3, 5, 5, 5},
			},
			expected: expected{
				result: -1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			majorityElement1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_majorityElement2(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 2, 3},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			expected: expected{
				result: 2,
			},
		},
		// {
		// 	name: "3",
		// 	args: args{
		// 		prices: []int{3, 1, 3, 4, 3, 4, 1, 4, 3, 3, 3, 3, 3, 5, 5, 5},
		// 	},
		// 	expected: expected{
		// 		result: -1,
		// 	},
		// },
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			majorityElement2(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_majorityElement3(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 2, 3},
			},
			expected: expected{
				result: 3,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			expected: expected{
				result: 2,
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 3, 4, 3, 4, 1, 4, 3, 3, 3, 3, 3, 5, 5, 5},
			},
			expected: expected{
				result: -1,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			majorityElement3(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
