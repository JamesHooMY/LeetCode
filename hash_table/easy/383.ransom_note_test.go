package easy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/ransom-note/description/

// method 1 hash table
// 1) use a hash table to store the characters and their counts in magazine
// 2) use a for loop to scan the characters in ransomNote
// 3) if the character is not in the hash table, then return false
// 4) if the character is in the hash table, then minus 1 from the count of the character in the hash table
// 5) if the count of the character in the hash table is less than 0, then return false
// 6) return true
func canConstruct1(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	magazineMap := make(map[rune]int) // key: character, value: count of the character
	for _, v := range magazine {
		magazineMap[v]++
	}

	for _, v := range ransomNote {
		// magazineMap[v] == 0, mean the character is not for magazineMap[v]--
		if magazineMap[v] == 0 {
			return false
		}
		magazineMap[v]--
	}

	return true
}

// method 2 array
// 1) use an array to store the counts of the characters in magazine
// 2) use a for loop to scan the characters in ransomNote
// 3) if the character is not in the array, then return false
// 4) if the character is in the array, then minus 1 from the count of the character in the array
// 5) if the count of the character in the array is less than 0, then return false
// 6) return true
// TC = O(N), SC = O(1)
// * this is the best solution for me currently
func canConstruct2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	magazineArray := [26]int{} // index: character - 'a', value: count of the character

	// for _, v := range magazine {
	// 	magazineArray[v-'a']++
	// }

	// for _, v := range ransomNote {
	// 	if magazineArray[v-'a'] == 0 {
	// 		return false
	// 	}
	// 	magazineArray[v-'a']--
	// }

	// * this performance is better than for range
	for i := 0; i < len(magazine); i++ {
		magazineArray[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		if magazineArray[ransomNote[i]-'a'] == 0 {
			return false
		}
		magazineArray[ransomNote[i]-'a']--
	}

	return true
}

func Test_canConstruct1(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
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
				ransomNote: "a",
				magazine:   "b",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			canConstruct1(tc.args.ransomNote, tc.args.magazine),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_canConstruct2(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
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
				ransomNote: "a",
				magazine:   "b",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			expected: expected{
				result: false,
			},
		},
		{
			name: "3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			expected: expected{
				result: true,
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			canConstruct2(tc.args.ransomNote, tc.args.magazine),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_canConstruct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct1("aa", "aab")
	}
}

func Benchmark_canConstruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canConstruct2("aa", "aab")
	}
}
