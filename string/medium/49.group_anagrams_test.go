package medium

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

// method 1 hash table
// 1) use a strSliceStrMap to store the sortedStr of the string
// 2) for each string in strs, sort the string, then check if the sortedStr is in the strSliceStrMap, if yes, then append the string to the value of the strSliceStrMap, if no, then add the sortedStr to the strSliceStrMap, and append the string to the value of the strSliceStrMap
// 3) finally, return the values of the strSliceStrMap
// TC = O(N * KlogK), SC = O(N * K), N is the length of strs, K is the average length of string in strs
// * this is the best solution for me currently
func groupAnagrams1(strs []string) [][]string {
	result := [][]string{}
	if len(strs) == 0 {
		return result
	}

	// use a strSliceStrMap to store the sortedStr of the string
	/*
		str -> sorted string

		"eat" -> "aet"
		"tea" -> "aet"
		"tan" -> "ant"
		"ate" -> "aet"
		"nat" -> "ant"
		"bat" -> "abt"
	*/
	strSliceStrMap := make(map[string][]string) // key: sorted string, value: []anagrams
	// * manipulate value of slice strs is faster than manipulate index of slice strs
	for _, str := range strs {
		sortedStr := sortString(str)
		strSliceStrMap[sortedStr] = append(strSliceStrMap[sortedStr], str)
	}
	// for i := range strs {
	// 	sortedStr := sortString(strs[i])
	// 	strSliceStrMap[sortedStr] = append(strSliceStrMap[sortedStr], strs[i])
	// }

	// * manipulate the values of strSliceStrMap is faster than manipulate the keys of strSliceStrMap
	for _, sliceStr := range strSliceStrMap {
		result = append(result, sliceStr)
	}
	// for k := range strSliceStrMap {
	// 	result = append(result, strSliceStrMap[k])
	// }

	return result
}

func sortString(str string) string {
	// convert string to slice
	sliceByte := []byte(str)

	// use sort.Slice to sort the slice, TC = O(NlogN), SC = O(logN)
	sort.Slice(sliceByte, func(i, j int) bool {
		return sliceByte[i] < sliceByte[j]
	})

	// convert slice to string
	return string(sliceByte)
}

func Test_groupAnagrams1(t *testing.T) {
	type args struct {
		strs []string
	}
	type expected struct {
		result [][]string
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
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			expected: expected{
				result: [][]string{
					{"eat", "tea", "ate"},
					{"tan", "nat"},
					{"bat"},
				},
			},
		},
		{
			name: "2",
			args: args{
				strs: []string{""},
			},
			expected: expected{
				result: [][]string{
					{""},
				},
			},
		},
		{
			name: "3",
			args: args{
				strs: []string{"a"},
			},
			expected: expected{
				result: [][]string{
					{"a"},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(
			t,
			tc.expected.result,
			groupAnagrams1(tc.args.strs),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_groupAnagrams1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		groupAnagrams1([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	}
}
