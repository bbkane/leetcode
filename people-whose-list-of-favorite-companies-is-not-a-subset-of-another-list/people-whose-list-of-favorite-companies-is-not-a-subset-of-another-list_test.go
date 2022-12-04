package peoplewhoselistoffavoritecompaniesisnotasubsetofanotherlist

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type Node struct {
	Children    map[string]*Node
	PersonIndex int
}

func addToRet(cur *Node, ret *[]int) {
	if cur.Children == nil {
		*ret = append(*ret, cur.PersonIndex)
	} else {
		for _, child := range cur.Children {
			addToRet(child, ret)
		}
	}

}

func peopleIndexes(favoriteCompanies [][]string) []int {
	root := &Node{
		Children:    nil,
		PersonIndex: 0,
	}
	// for companies in favoriteCompanies:
	for personIndex, companies := range favoriteCompanies {
		cur := root
		// sort companies
		sort.Strings(companies)
		for _, company := range companies {
			// insert personIndex, company into tree, making nodes if needed.
			if cur.Children == nil {
				cur.Children = make(map[string]*Node)
			}
			child, exists := cur.Children[company]
			if exists {
				cur = child
			} else {
				cur.Children[company] = &Node{
					Children:    nil,
					PersonIndex: 0,
				}
				cur = cur.Children[company]
			}
		}
		cur.PersonIndex = personIndex
	}

	// the leaves of the tree are the not-subset indexes
	ret := []int{}
	addToRet(root, &ret)
	sort.Ints(ret)
	return ret
}

func Test_peopleIndexes(t *testing.T) {
	type args struct {
		favoriteCompanies [][]string
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "1",
			args: args{
				favoriteCompanies: [][]string{
					{"leetcode", "google", "facebook"},
					{"google", "microsoft"},
					{"google", "facebook"},
					{"google"},
					{"amazon"},
				},
			},
			expected: []int{0, 1, 4},
		},
		{
			name: "2",
			args: args{
				favoriteCompanies: [][]string{
					{"leetcode", "google", "facebook"},
					{"leetcode", "amazon"},
					{"facebook", "google"},
				},
			},
			expected: []int{0, 1},
		},
		{
			name: "3",
			args: args{
				favoriteCompanies: [][]string{
					{"leetcode"},
					{"google"},
					{"facebook"},
					{"amazon"},
				},
			},
			expected: []int{0, 1, 2, 3},
		},
		{
			// simplest case that breaks the tree
			name: "11-simple",
			args: args{
				favoriteCompanies: [][]string{
					{"a", "b", "c"},
					{"a", "c"},
				},
			},
			expected: []int{0},
		},
		{
			name: "11",
			args: args{
				favoriteCompanies: [][]string{
					{
						"nxaqhyoprhlhvhyojanr",
						"ovqdyfqmlpxapbjwtssm",
						"qmsbphxzmnvflrwyvxlc",
						"udfuxjdxkxwqnqvgjjsp",
						"yawoixzhsdkaaauramvg",
						"zycidpyopumzgdpamnty",
					},
					{
						"nxaqhyoprhlhvhyojanr",
						"ovqdyfqmlpxapbjwtssm",
						"udfuxjdxkxwqnqvgjjsp",
						"yawoixzhsdkaaauramvg",
						"zycidpyopumzgdpamnty",
					},
					{
						"ovqdyfqmlpxapbjwtssm",
						"qmsbphxzmnvflrwyvxlc",
						"udfuxjdxkxwqnqvgjjsp",
						"yawoixzhsdkaaauramvg",
						"zycidpyopumzgdpamnty",
					},
				},
			},
			expected: []int{0},
		},
		{
			name: "11a",
			args: args{
				favoriteCompanies: [][]string{
					{
						"n",
						"o",
						"q",
						"u",
						"y",
						"z",
					},
					{
						"n",
						"o",
						"u",
						"y",
						"z",
					},
					{
						"o",
						"q",
						"u",
						"y",
						"z",
					},
				},
			},
			expected: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := peopleIndexes(tt.args.favoriteCompanies)
			require.Equal(t, tt.expected, actual)
		})
	}
}
