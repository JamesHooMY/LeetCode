package medium

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-all-anagrams-in-a-string/

// method 1 hash table
// 1) use a strMap to store the sortedStr of the string
// 2) for each string in strs, sort the string, then check if the sortedStr is in the strMap, if yes, then append the string to the value of the strMap, if no, then add the sortedStr to the strMap, and append the string to the value of the strMap
// 3) finally, return the values of the strMap
// TC = O(N * KlogK), SC = O(N * K), N is the length of strs, K is the average length of string in strs
// * this is the best solution for me currently
func groupAnagrams1(strs []string) [][]string {
	result := [][]string{}
	if len(strs) == 0 {
		return result
	}

	// use a strMap to store the sortedStr of the string
	/*
		str -> sorted string

		"eat" -> "aet"
		"tea" -> "aet"
		"tan" -> "ant"
		"ate" -> "aet"
		"nat" -> "ant"
		"bat" -> "abt"
	*/
	strMap := make(map[string][]string) // key: sorted string, value: []anagrams
	// * manipulate value of slice strs is faster than manipulate index of slice strs
	for _, str := range strs {
		sortedStr := sortString(str)
		strMap[sortedStr] = append(strMap[sortedStr], str)
	}
	// for i := range strs {
	// 	sortedStr := sortString(strs[i])
	// 	strMap[sortedStr] = append(strMap[sortedStr], strs[i])
	// }

	// * manipulate the values of strMap is faster than manipulate the keys of strMap
	for _, v := range strMap {
		result = append(result, v)
	}
	// for k := range strMap {
	// 	result = append(result, strMap[k])
	// }

	return result
}

func sortString(str string) string {
	// convert string to slice
	strSlice := []byte(str)

	// use sort.Slice to sort the slice, TC = O(NlogN), SC = O(logN)
	sort.Slice(strSlice, func(i, j int) bool {
		return strSlice[i] < strSlice[j]
	})

	// convert slice to string
	return string(strSlice)
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
