package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/top-k-frequent-elements/submissions/1331210666/

// method 1 hash table + sort
// 1) use a hash table to store the numbers and their counts
// 2) use a for loop to scan the numbers in nums
// 3) store the number and its count in the hash table
// 4) use a for loop to scan the keys in the hash table
// 5) store the keys in a slice
// 6) sort the slice in descending order
// 7) return the first k elements in the slice
// TC = O(NlogN), SC = O(N)
func topKFrequent1(nums []int, k int) []int {
	numMap := map[int]int{} // key: number, value: count of the number

	for _, num := range nums {
		numMap[num]++
	}

	keys := make([]int, 0, len(numMap))

	for key := range numMap {
		keys = append(keys, key)
	}

	// descending order
	sort.Slice(keys, func(i, j int) bool {
		return numMap[keys[i]] > numMap[keys[j]]
	})

	return keys[:k]
}

// method 2 hash table + bucket sort
// 1) use a hash table to store the numbers and their counts
// 2) use a for loop to scan the numbers in nums
// 3) store the number and its count in the hash table
// 4) use a bucket to store the numbers
// 5) the index of the bucket is the count of the number
// 6) the value of the bucket is a slice of the numbers
// 7) use a for loop to scan the keys in the hash table
// 8) store the numbers in the bucket
// 9) use a for loop to scan the bucket in descending order
// 10) store the numbers in the result slice
// 11) return the result slice
// TC = O(N), SC = O(N)
// * this is the best solution for me currently
func topKFrequent2(nums []int, k int) []int {
	numMap := map[int]int{}

	for _, num := range nums {
		numMap[num]++
	}

	// len(nums)+1, because the index is count
	// [1,2,3] -> the total count is 3, so the bucket length is 4, then the index is 0, 1, 2, 3
	bucket := make([][]int, len(nums)+1) // index: count, value: []int of the numbers
	for num, count := range numMap {
		bucket[count] = append(bucket[count], num)
	}

	result := make([]int, 0, k)
	for i := len(nums); i >= 0 && len(result) < k; i-- {
		if len(result) == k {
			break
		}

		for _, num := range bucket[i] {
			if len(result) == k {
				break
			}
			result = append(result, num)
		}
	}

	return result
}

func Test_topKFrequent1(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			expected: expected{
				result: []int{1, 2},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 1, 2, 100},
				k:    3,
			},
			expected: expected{
				result: []int{1, 2, 3},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, topKFrequent1(tc.args.nums, tc.args.k), tc.name)
	}
}

func Test_topKFrequent2(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			expected: expected{
				result: []int{1, 2},
			},
		},
		{
			name: "2",
			args: args{
				nums: []int{1},
				k:    1,
			},
			expected: expected{
				result: []int{1},
			},
		},
		{
			name: "3",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 1, 2, 100},
				k:    3,
			},
			expected: expected{
				result: []int{1, 2, 3},
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected.result, topKFrequent2(tc.args.nums, tc.args.k), tc.name)
	}
}

// benchmark
func Benchmark_topKFrequent1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent1([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 1, 2, 100}, 3)
	}
}

func Benchmark_topKFrequent2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topKFrequent2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 1, 1, 2, 100}, 3)
	}
}
