package leetcode_2

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	     Val int
	     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum1 := GetSum(l1)
	sum2 := GetSum(l2)
	return GetNode(sum1+sum2)
}


func GetSum(l *ListNode)  int {
	sum := 0
	i := 0
	pre := l
	for pre.Next != nil {
		sum = sum + pre.Val * pow(10,i)
		pre = pre.Next
		i++
	}

	return  sum + pre.Val * pow(10,i)
}

func GetNode(sum int) *ListNode  {
	arr := []ListNode{}
	if sum == 0 {
		return &ListNode{0,nil}
	}
	for sum > 0{
		arr = append(arr,ListNode{sum%10,nil})
		sum = int(sum/10)
	}


	for index := range arr {
		if index >0 {
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
