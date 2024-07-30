package codility

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
	1) replace the '?' with 'a' or 'b'
	2) make sure filtered string not contain three consecutive characters like 'aaa' or 'bbb'
*/

// method 1 brute force
// 1) convert the string to byte array
// 2) use for loop to scan the byte array
// 3) if the current character is '?', then iterate the options 'a' and 'b'
// 4) if the current character is the first character, then replace the different character with the next character
// 5) if the current character is the last character, then replace the different character with the previous character
// 6) if the current character is in the middle, if the left and right characters are the same with opt, then skip, otherwise replace the current character with the opt
// 7) if the current character is in the middle, if the left and right characters are not same, then start leftScan and rightScan
// 8) leftScan: scan the left characters, if the count of the same character is less than 2, then replace the current character with the opt
// 9) rightScan: scan the right characters, if the count of the same character is less than 2, then replace the current character with the opt
// 10) return the filtered string
// TC: O(N^2), SC: O(N)
func filteredString1(s string) string {
	filteredString := []byte(s)
	options := []byte{'a', 'b'}

	for i := 0; i < len(filteredString); i++ {
		if filteredString[i] == '?' {
			for _, opt := range options {
				// first '?'
				if i == 0 && opt != filteredString[i+1] {
					filteredString[i] = opt
					break
				}

				// last '?'
				if i == len(filteredString)-1 && opt != filteredString[i-1] {
					filteredString[i] = opt
					break
				}

				if i > 0 && i < len(filteredString)-1 {
					// "a?a", "a" --> "a?a", "b"
					if opt == filteredString[i-1] && opt == filteredString[i+1] {
						continue
					}

					// "a?a", "b" --> "aba"
					if opt != filteredString[i-1] && opt != filteredString[i+1] {
						filteredString[i] = opt
						break
					}

					// left and right characters are not the same ---------------------
					// check left
					if leftScan(i-1, opt, filteredString) {
						filteredString[i] = opt
						break
					}

					// check right
					if rightScan(i+1, opt, filteredString) {
						filteredString[i] = opt
						break
					}
					// ----------------------------------------------------------------
				}
			}
		}
	}

	return string(filteredString)
}

func leftScan(i int, char byte, filteredString []byte) bool {
	count := 0
	for j := i; j >= 0; j-- {
		if filteredString[j] == char {
			count++
		}
	}

	return count < 2
}

func rightScan(i int, char byte, filteredString []byte) bool {
	count := 0
	for j := i; j < len(filteredString); j++ {
		if filteredString[j] == char {
			count++
		}
	}

	return count < 2
}

// method 2 hash table, always keep the count between min 0 and max 2
// 1) convert the string to byte array, and create the options array, and hash table to store the count of the options
// 2) use for loop to scan the byte array
// 3) if the current character is '?', then iterate the options 'a' and 'b'
// 4) if the current character is the first character, then replace the different character with the next character, and increase the count of the opt
// 5) if the current character is the last character, then replace the different character with the previous character
// 6) if the current character is in the middle, if the left and right characters are not same, left character is the same with opt and the count of the opt is less than 2, then replace the current character with the opt, and increase the count of the opt, and decrease the count of the right character
// 7) if the current character is in the middle, if the left and right characters are the same with opt, then skip, otherwise replace the current character with the opt, and increase the count of the opt, and decrease the count of the left or right character
// 8) return the filtered string
// TC: O(N), SC: O(N), slower than method 1 !!!
func filteredString2(s string) string {
	filteredString := []byte(s)
	options := [2]byte{'a', 'b'}
	// hash table
	optionsMap := map[byte]int{
		'a': 0, // keep min 0 and max 2
		'b': 0, // keep min 0 and max 2
	}

	for i := 0; i < len(filteredString); i++ {
		left, right := i-1, i+1
		if filteredString[i] == '?' {
			for _, opt := range options {
				// first '?'
				if i == 0 && opt != filteredString[right] {
					filteredString[i] = opt
					optionsMap[opt]++
					break
				}

				// last '?'
				if i == len(filteredString)-1 && opt != filteredString[left] {
					filteredString[i] = opt
					break
				}

				// middle '?'
				if i > 0 && i < len(filteredString)-1 {
					// "aa?b", "b" --> "aabb"
					if filteredString[left] != filteredString[right] && opt == filteredString[left] && optionsMap[opt] < 2 {
						filteredString[i] = opt
						optionsMap[opt]++
						optionsMap[filteredString[right]]--
						break
					}

					// "aa?a", "a" --> "aa?a", "b"
					if opt == filteredString[left] && opt == filteredString[right] {
						continue
					}

					// "aa?a", "b" --> "aaba"
					if opt != filteredString[left] && opt != filteredString[right] {
						filteredString[i] = opt
						if optionsMap[opt] < 2 {
							optionsMap[opt]++
						}
						if optionsMap[filteredString[left]] > 0 {
							optionsMap[filteredString[left]]--
						}
						break
					}
				}
			}
		} else {
			optionsMap[filteredString[i]]++

			if i > 0 && i < len(filteredString)-1 {
				if filteredString[i] != filteredString[left] && optionsMap[filteredString[left]] > 0 {
					optionsMap[filteredString[left]]--
				}
			}
		}
	}

	return string(filteredString)
}

// method 3 slice counts (*slice is faster than hash table !!!), always keep the count between min 0 and max 2
// 1) convert the string to byte array, and create the options array, and slice to store the count of the options
// 2) use for loop to scan the byte array
// 3) if the current character is '?', then iterate the options 'a' and 'b'
// 4) if the current character is the first character, then replace the different character with the next character, and increase the count of the opt
// 5) if the current character is the last character, then replace the different character with the previous character
// 6) if the current character is in the middle, if the left and right characters are not same, left character is the same with opt and the count of the opt is less than 2, then replace the current character with the opt, and increase the count of the opt, and decrease the count of the right character
// 7) if the current character is in the middle, if the left and right characters are the same with opt, then skip, otherwise replace the current character with the opt, and increase the count of the opt, and decrease the count of the left or right character
// 8) return the filtered string
// TC: O(N), SC: O(N)
// * this is the best solution for me currently
func filteredString3(s string) string {
	filteredString := []byte(s)
	options := [2]byte{'a', 'b'}
	// index 0: 'a', index 1: 'b'
	optionsArr := [2]int{0, 0} // keep min 0 and max 2

	for i := 0; i < len(filteredString); i++ {
		left, right := i-1, i+1
		if filteredString[i] == '?' {
			for _, opt := range options {
				// first '?'
				if i == 0 && opt != filteredString[right] {
					filteredString[i] = opt
					optionsArr[opt-'a']++
					break
				}

				// last '?'
				if i == len(filteredString)-1 && opt != filteredString[left] {
					filteredString[i] = opt
					break
				}

				// middle '?'
				if i > 0 && i < len(filteredString)-1 {
					// "aa?b", "b" --> "aabb"
					if filteredString[left] != filteredString[right] && opt == filteredString[left] && optionsArr[opt-'a'] < 2 {
						filteredString[i] = opt
						optionsArr[opt-'a']++
						optionsArr[filteredString[right]-'a']--
						break
					}

					// "aa?a", "a" --> "aa?a", "b"
					if opt == filteredString[left] && opt == filteredString[right] {
						continue
					}

					// "aa?a", "b" --> "aaba"
					if opt != filteredString[left] && opt != filteredString[right] {
						filteredString[i] = opt
						if optionsArr[opt-'a'] < 2 {
							optionsArr[opt-'a']++
						}
						if optionsArr[filteredString[left]-'a'] > 0 {
							optionsArr[filteredString[left]-'a']--
						}
						break
					}
				}
			}
		} else {
			optionsArr[filteredString[i]-'a']++

			if i > 0 && i < len(filteredString)-1 {
				if filteredString[i] != filteredString[left] && optionsArr[filteredString[left]-'a'] > 0 {
					optionsArr[filteredString[left]-'a']--
				}
			}
		}
	}

	return string(filteredString)
}

func Test_filteredString1(t *testing.T) {
	type args struct {
		s string
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
				s: "a?bb",
			},
			expected: expected{
				result: "aabb",
			},
		},
		{
			name: "2",
			args: args{
				s: "??abb",
			},
			expected: expected{
				result: "ababb",
			},
		},
		{
			name: "3",
			args: args{
				s: "a?b?aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "4",
			args: args{
				s: "aa??aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "5",
			args: args{
				s: "?ab?",
			},
			expected: expected{
				result: "baba",
			},
		},
		{
			name: "6",
			args: args{
				s: "a?b?aab?bab??a??baa??????a?b?a",
			},
			expected: expected{
				result: "aabbaababababababaabababbaabaa",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			filteredString1(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_filteredString2(t *testing.T) {
	type args struct {
		s string
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
				s: "a?bb",
			},
			expected: expected{
				result: "aabb",
			},
		},
		{
			name: "2",
			args: args{
				s: "??abb",
			},
			expected: expected{
				result: "ababb",
			},
		},
		{
			name: "3",
			args: args{
				s: "a?b?aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "4",
			args: args{
				s: "aa??aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "5",
			args: args{
				s: "?ab?",
			},
			expected: expected{
				result: "baba",
			},
		},
		{
			name: "6",
			args: args{
				s: "a?b?aab?bab??a??baa??????a?b?a",
			},
			expected: expected{
				result: "aabbaababababababaabababbaabba",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			filteredString2(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

func Test_filteredString3(t *testing.T) {
	type args struct {
		s string
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
				s: "a?bb",
			},
			expected: expected{
				result: "aabb",
			},
		},
		{
			name: "2",
			args: args{
				s: "??abb",
			},
			expected: expected{
				result: "ababb",
			},
		},
		{
			name: "3",
			args: args{
				s: "a?b?aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "4",
			args: args{
				s: "aa??aa",
			},
			expected: expected{
				result: "aabbaa",
			},
		},
		{
			name: "5",
			args: args{
				s: "?ab?",
			},
			expected: expected{
				result: "baba",
			},
		},
		{
			name: "6",
			args: args{
				s: "a?b?aab?bab??a??baa??????a?b?a",
			},
			expected: expected{
				result: "aabbaababababababaabababbaabba",
			},
		},
	}

	for _, tc := range testCases {
		assert.Equal(
			t,
			tc.expected.result,
			filteredString3(tc.args.s),
			fmt.Sprintf("testCase name: %s", tc.name),
		)
	}
}

// benchmark
func Benchmark_filteredString1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filteredString1("a?b?aab?bab??a??baa??????a?b?a")
	}
}

func Benchmark_filteredString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filteredString2("a?b?aab?bab??a??baa??????a?b?a")
	}
}

func Benchmark_filteredString3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		filteredString3("a?b?aab?bab??a??baa??????a?b?a")
	}
}