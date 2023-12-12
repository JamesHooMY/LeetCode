package medium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/longest-consecutive-sequence/description/

// method 1
// 1) use a map to store all nums without duplicate
// 2) loop the map, find the start of the sequence, and accumulate the length of the sequence into currentLength
// 3) compare currentLength with maxLength, if currentLength > maxLength, update maxLength
// 4) return maxLength
// TC = O(n), SC = O(n)
// * this is the best solution for me currently
func longestConsecutive1(nums []int) int {
	numMap := make(map[int]bool) // SC = O(n)
	maxLength := 0

	// put all nums without duplicate into map
	for _, num := range nums { // TC = O(n)
		numMap[num] = true
	}

	for num := range numMap { // TC = O(n)
		// this condition make sure the num is the start of the sequence
		// if num-1 is in the map, it means the sequence will be start from num-1
		if !numMap[num-1] {
			currentNum := num
			currentLength := 1

			// find the end of the sequence
			// this step will not cause the "for num := range numMap" from O(N) to O(N^2), due to the numbers in the numMap are unique !!!
			for numMap[currentNum+1] {
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

func Test_longestConsecutive1(t *testing.T) {
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
				nums: []int{100, 4, 200, 1, 3, 2},
			},
			expected: expected{
				result: 4,
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			},
			expected: expected{
				result: 9,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			longestConsecutive1(tc.args.nums),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}
