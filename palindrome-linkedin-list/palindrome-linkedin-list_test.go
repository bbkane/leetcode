package palindromelinkedinlist

// https://leetcode.com/problems/palindrome-linked-list/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// get the length of the list
	listLen := 0
	for cur := head; cur != nil; cur = cur.Next {
		listLen += 1
	}

	// make an array of len/2
	arr := make([]int, listLen/2)

	// go to the halfway point adding to the array
	cur := head
	for i := 0; i < listLen/2; i++ {
		arr[i] = cur.Val
		cur = cur.Next
	}
	// skip the middle element if needed
	if listLen%2 != 0 {
		cur = cur.Next
	}

	// go through th rest of the list and subtract
	for i := listLen/2 - 1; i >= 0; i-- {
		arr[i] -= cur.Val
		cur = cur.Next
	}

	// check if all elements in the array are 0
	for _, el := range arr {
		if el != 0 {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name        string
		head        *ListNode
		expectedRes bool
	}{
		{
			name:        "[1,2,2,1]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}},
			expectedRes: true,
		},
		{
			name:        "[1,2]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expectedRes: false,
		},
		{

			name:        "nil",
			head:        nil,
			expectedRes: true,
		},
		{
			name:        "[1]",
			head:        &ListNode{Val: 1, Next: nil},
			expectedRes: true,
		},
		{
			name:        "[1,3,1]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 1, Next: nil}}},
			expectedRes: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRes := isPalindrome(tt.head)
			require.Equal(t, tt.expectedRes, actualRes)
		})
	}
}
