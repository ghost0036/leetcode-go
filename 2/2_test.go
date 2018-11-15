package leetcode_2

import (
	"fmt"
	"testing"
)

func TestNomal(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	//s := addTwoNumbers(l1, l2)
	s := addTwoNumbers1(l1, l2)
	fmt.Println(s.Val, s.Next.Val, s.Next.Next.Val)
}

func TestException(t *testing.T) {
	l1 := &ListNode{1, &ListNode{8, nil}}
	l2 := &ListNode{0, nil}

	s := addTwoNumbers1(l1, l2)
	fmt.Println(s.Val, s.Next.Val)
}

func TestException1(t *testing.T) {
	l1 := &ListNode{5, nil}
	l2 := &ListNode{5, nil}

	s := addTwoNumbers1(l1, l2)
	fmt.Println(s.Val, s.Next.Val)
}

func TestException2(t *testing.T) {
	l1 := &ListNode{1, nil}
	l2 := &ListNode{9, &ListNode{9, nil}}

	s := addTwoNumbers1(l1, l2)
	fmt.Println(s.Val, s.Next.Val, s.Next.Next.Val)
}
