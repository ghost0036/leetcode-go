package leetcode_19

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	before := head

	for n != 0 {
		n = n - 1
		cur = cur.Next
	}

	if cur == nil {
		return head.Next
	}

	for cur.Next != nil {
		cur = cur.Next
		before = before.Next
	}

	cur.Next = cur.Next.Next
	return head
}
