package addtwonumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Ptr := l1
	l2Ptr := l2
	prevSumPtr := &ListNode{Val: 0, Next: nil}
	startPtr := prevSumPtr
	carry := 0
	for l1Ptr != nil || l2Ptr != nil {
		l1Addend := 0
		if l1Ptr != nil {
			l1Addend = l1Ptr.Val
			l1Ptr = l1Ptr.Next
		}
		l2Addend := 0
		if l2Ptr != nil {
			l2Addend = l2Ptr.Val
			l2Ptr = l2Ptr.Next
		}
		sum := l1Addend + l2Addend + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		prevSumPtr.Next = &ListNode{Val: sum, Next: nil}
		prevSumPtr = prevSumPtr.Next
	}
	if carry == 1 {
		prevSumPtr.Next = &ListNode{Val: 1, Next: nil}
	}
	return startPtr.Next
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name     string
		args     args
		expected *ListNode
	}{
		{
			name: "carry",
			args: args{
				l1: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
				l2: &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			},
			expected: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
		{
			name: "carry",
			args: args{
				l1: &ListNode{Val: 9, Next: nil},
				l2: &ListNode{Val: 1, Next: nil},
			},
			expected: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}},
		},
		{
			name: "b2",
			args: args{
				l1: &ListNode{Val: 1, Next: nil},
				l2: &ListNode{Val: 2, Next: nil},
			},
			expected: &ListNode{Val: 3, Next: nil},
		},
		{
			name: "1",
			args: args{
				l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			},
			expected: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
		{
			name: "2",
			args: args{
				l1: &ListNode{Val: 0, Next: nil},
				l2: &ListNode{Val: 0, Next: nil},
			},
			expected: &ListNode{Val: 0, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.expected) {
			// 	t.Errorf("addTwoNumbers() = %v, want %v", got, tt.expected)
			// }
			actual := addTwoNumbers(tt.args.l1, tt.args.l2)
			require.Equal(t, tt.expected, actual)
		})
	}
}
