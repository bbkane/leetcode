package middleofthelinkedlist

// https://leetcode.com/problems/middle-of-the-linked-list/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	listLen := 0
	for cur := head; cur != nil; cur = cur.Next {
		listLen += 1
	}
	// go to the halfway - 1 point
	cur := head
	for i := 0; i < listLen/2-1; i++ {
		cur = cur.Next
	}

	cur = cur.Next
	return cur
}

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name        string
		head        *ListNode
		expectedRes *ListNode
	}{
		{
			name:        "[1,2,3,4,5]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
			expectedRes: &ListNode{Val: 3},
		},
		{
			name:        "[1,2,3,4,5,6]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}}}}},
			expectedRes: &ListNode{Val: 4},
		},
		{
			name:        "nil",
			head:        nil,
			expectedRes: nil,
		},
		{
			name:        "[1]",
			head:        &ListNode{Val: 1, Next: nil},
			expectedRes: &ListNode{Val: 1, Next: nil},
		},
		{
			name:        "[1,2]",
			head:        &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expectedRes: &ListNode{Val: 2, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualRes := middleNode(tt.head)
			if tt.expectedRes == nil && actualRes == nil {
				return
			}
			require.NotNil(t, tt.expectedRes, "expectedRes")
			require.NotNil(t, actualRes, "actualRes")
			// I'm gonna assume the val is unique for the purpose of tests
			require.Equal(t, tt.expectedRes.Val, actualRes.Val)
		})
	}
}
