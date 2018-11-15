package leetcode_2

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

/**
错误解答；无法解决溢出问题
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum1 := GetSum(l1)
	sum2 := GetSum(l2)
	return GetNode(sum1 + sum2)
}

/**
通过
*/
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	cur := result

	list1 := l1
	list2 := l2
	isNeed := 0
	for list1 != nil || list2 != nil {
		var node ListNode
		v1 := 0
		v2 := 0

		if list1 != nil {
			v1 = list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			v2 = list2.Val
			list2 = list2.Next
		}

		sum := v1 + v2 + isNeed

		if sum > 9 {
			isNeed = 1
		} else {
			isNeed = 0
		}
		node.Val = sum % 10

		cur.Next = &node
		cur = &node
	}

	if isNeed == 1 {
		cur.Next = &ListNode{1, nil}
	}

	//if cur.Val >= 10 {
	//	cur.Val = cur.Val % 10
	//	cur.Next = &ListNode{1,nil}
	//}

	return result.Next
}

func GetSum(l *ListNode) int {
	sum := 0
	i := 0
	pre := l
	for pre.Next != nil {
		sum = sum + pre.Val*pow(10, i)
		pre = pre.Next
		i++
	}

	return sum + pre.Val*pow(10, i)
}

func GetNode(sum int) *ListNode {
	arr := []ListNode{}
	if sum == 0 {
		return &ListNode{0, nil}
	}
	for sum > 0 {
		arr = append(arr, ListNode{sum % 10, nil})
		sum = int(sum / 10)
	}

	for index := range arr {
		if index > 0 {
			arr[index-1].Next = &arr[index]
		}
	}
	return &arr[0]

}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}
