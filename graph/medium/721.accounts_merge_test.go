package medium

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/accounts-merge/description/

// method 1 union find
// 1) use a union find to store the parent of each email
// 2) use a hash table to store the email to name mapping
// 3) iterate the accounts
// 4) iterate the emails of the account
// 5) if the email is not in the union find, add it to the union find
// 6) store the email to name mapping
// 7) if the index is greater than 1, union the first email with the other emails
// 8) use a hash table to store the root to emails mapping
// 9) iterate the email to name mapping
// 10) find the root of the email
// 11) store the email to the root mapping
// 12) store the root to email mapping
// 13) iterate the root to email mapping
// 14) sort the emails
// 15) append the account to the result
// 16) return the result
// TC: O(NlogN), SC: O(N), where N is the number of emails
func accountsMerge1(accounts [][]string) [][]string {
	uf := UnionFind[string]{
		parent: map[string]string{},
		rank:   map[string]int{},
	}
	emailNameMap := map[string]string{} // key: email, value: name

	for _, account := range accounts {
		name := account[0]

		// iterate the emails of the account
		for i := 1; i < len(account); i++ {
			if _, exist := uf.parent[account[i]]; !exist {
				uf.parent[account[i]] = account[i]
			}

			emailNameMap[account[i]] = name

			if i > 1 {
				uf.Union(account[1], account[i]) // union the first email with the other emails
			}
		}
	}

	rootEmailsMap := map[string][]string{} // key: root, value: emails
	for email := range emailNameMap {
		root := uf.Find(email)
		rootEmailsMap[root] = append(rootEmailsMap[root], email)
	}

	result := [][]string{}
	for root, emails := range rootEmailsMap {
		sort.Strings(emails) // TC: O(NlogN), SC: O(N)
		account := append([]string{emailNameMap[root]}, emails...)
		result = append(result, account)
	}

	return result
}

type UnionFind[T comparable] struct {
	parent map[T]T   // key: child, value: parent
	rank   map[T]int // key: parent, value: rank, rank is the depth of the tree rooted at parent i
}

// TC: O(N) where N is the number of emails， SC: O(N) where N is the depth of the tree
func (uf *UnionFind[T]) Find(x T) (root T) {
	// * if the parent of x is not x, then recursively find the parent of x, all elements have the same parent root
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

// TC: O(N) where N is the number of emails， SC: O(N) where N is the depth of the tree
func (uf *UnionFind[T]) Union(x, y T) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

// method 2 union find
// 1) use a parents to store the parent of each account
// 2) use a emailIdxMap to store the index of the parent
// 3) iterate the accounts
// 4) iterate the emails of the account
// 5) if the email is in the emailIdxMap, union the parent of the account with the parent of the email
// 6) store the email to index mapping
// 7) use a parentEmailsMap to store the parent to emails mapping
// 8) iterate the email to index mapping
// 9) find the root of the parent
// 10) append the email to the parent
// 11) sort the emails
// 12) append the account to the result
// 13) return the result
// TC: O(NlogN), SC: O(N), where N is the number of emails
// * this is the best solution for me currently
func accountsMerge2(accounts [][]string) [][]string {
	parents := make([]int, len(accounts)) // index: account, value: parent of the account
	emailIdxMap := map[string]int{}       // key: email, value: index of the parent

	for i, account := range accounts {
		parents[i] = i

		for j := 1; j < len(account); j++ {
			email := account[j]

			if idx, exist := emailIdxMap[email]; exist {
				/*
					idx: parent of the account, value: index of the parent

					account1 = ["john", "email1", "email2"]
					account2 = ["john", "email1", "email3"]
					account3 = ["jonh", "email3", "email4"]

					parents = [0, 1] -> [0, 0]
					parents = [0, 0, 2] -> [0, 0, 0]
				*/
				parents[find(parents, i)] = find(parents, idx)
			} else {
				/*
					emailIdxMap = {
						"email1": 0,
						"email2": 0,
						"email3": 1,
						"email4": 2,
					}
				*/
				emailIdxMap[email] = i
			}
		}
	}

	parentEmailsMap := map[int][]string{} // key: parent, value: emails
	for email, idx := range emailIdxMap {
		root := find(parents, idx)
		parentEmailsMap[root] = append(parentEmailsMap[root], email)
	}

	result := [][]string{}
	for parent, emails := range parentEmailsMap {
		sort.Strings(emails)
		account := append([]string{accounts[parent][0]}, emails...)
		result = append(result, account)
	}

	return result
}

func find(parents []int, i int) int {
	if parents[i] != i {
		parents[i] = find(parents, parents[i])
	}

	return parents[i]
}

func Test_accountsMerge1(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	type expected struct {
		result [][]string
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"Mary", "mary@mail.com"},
					{"John", "johnnybravo@mail.com"},
				},
			},
			expected: expected{
				result: [][]string{
					{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"Mary", "mary@mail.com"},
				},
			},
		},
		{
			name: "2",
			args: args{
				accounts: [][]string{
					{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
					{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
					{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
					{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
					{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
				},
			},
			expected: expected{
				result: [][]string{
					{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
					{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
					{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
					{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
					{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, accountsMerge1(tc.args.accounts), tc.name)
	}
}

func Test_accountsMerge2(t *testing.T) {
	type args struct {
		accounts [][]string
	}
	type expected struct {
		result [][]string
	}
	testCases := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "1",
			args: args{
				accounts: [][]string{
					{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
					{"John", "johnsmith@mail.com", "john00@mail.com"},
					{"Mary", "mary@mail.com"},
					{"John", "johnnybravo@mail.com"},
				},
			},
			expected: expected{
				result: [][]string{
					{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
					{"John", "johnnybravo@mail.com"},
					{"Mary", "mary@mail.com"},
				},
			},
		},
		{
			name: "2",
			args: args{
				accounts: [][]string{
					{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
					{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
					{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
					{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
					{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
				},
			},
			expected: expected{
				result: [][]string{
					{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
					{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
					{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
					{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
					{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
				},
			},
		},
	}

	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.expected.result, accountsMerge2(tc.args.accounts), tc.name)
	}
}

// benchmark
func Benchmark_accountsMerge1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		accountsMerge1([][]string{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
			{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
			{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
			{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
		})
	}
}

func Benchmark_accountsMerge2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		accountsMerge2([][]string{
			{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
			{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
			{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
			{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
			{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
		})
	}
}
