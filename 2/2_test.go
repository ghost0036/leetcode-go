package leetcode_2

import (
	"fmt"
	"testing"
)

func TestNomal(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	s := addTwoNumbers(l1, l2)
	fmt.Println(s.Val, s.Next.Val, s.Next.Val)
}

func TestException(t *testing.T) {
	l1 := &ListNode{0, nil}
	l2 := &ListNode{0, nil}

	s := addTwoNumbers(l1, l2)
	fmt.Println(s.Val, s.Next.Val, s.Next.Val)
}
