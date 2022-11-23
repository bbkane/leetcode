package addtwonumbers

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/add-two-numbers/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Ptr := l1
	l2Ptr := l2
	sumPtr := &ListNode{Val: 0, Next: nil}
	ret := sumPtr
	carry := 0
	for l1Ptr.Next != nil || l2Ptr.Next != nil {
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
			sum = 10 - sum
			carry = 1
		} else {
			carry = 0
		}
		sumPtr.Val = sum
		sumPtr.Next = &ListNode{Val: 0, Next: nil}
		// TODO: advance ptr to next; also use prevSumPtr instead of sumptr
	}
	return ret
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				l1: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
				l2: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			},
			want: &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		},
		{
			name: "2",
			args: args{
				l1: &ListNode{Val: 0, Next: nil},
				l2: &ListNode{Val: 0, Next: nil},
			},
			want: &ListNode{Val: 0, Next: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
