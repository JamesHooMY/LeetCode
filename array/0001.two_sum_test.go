package array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// method 1
// 1) use two for loop
// 2) The first loop will be used to scan each number in the target minus nums, while the second loop will confirm whether each result is equal to the following number. If they are equal, we have found two targets.
// TC = O(N^2), SC = (O)1
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		result := target - nums[i]

		for j := i + 1; j < len(nums); j++ {
			if result == nums[j] {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// method 2
// 1) use one for loop and one map
// 2) when running the for loop, use one resultMap (map[int]int, key is "result", value is "index") to store the result from target minus each num in nums
// 3) during the for loop scanning process, check each "result" and make sure whether it have been store in the map, if it was found in the map that mean we found the two targets
// TC = O(N), SC = O(N)
func twoSum2(nums []int, target int) []int {
	resultMap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		index, ok := resultMap[nums[i]]

		if ok {
			return []int{index, i}
		}

		resultMap[target-nums[i]] = i
	}

	return []int{}
}

func Test_twoSum1(t *testing.T) {
	type args struct {
		nums     []int
		target   int
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
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: expected{
				result: []int{0, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: expected{
				result: []int{1, 2},
			},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: expected{
				result: []int{0, 1},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			twoSum1(tc.args.nums, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_twoSum2(t *testing.T) {
	type args struct {
		nums     []int
		target   int
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
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			expected: expected{
				result: []int{0, 1},
			},
		},
		{
			name: "2",
			args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			expected: expected{
				result: []int{1, 2},
			},
		},
		{
			name: "3",
			args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			expected: expected{
				result: []int{0, 1},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			twoSum2(tc.args.nums, tc.args.target),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
