package leetcode_53

import "fmt"

func maxSubArray(nums []int) int {
	sum := nums[0] //记录当前总和
	max := nums[0] //记录最大值
	for _, v := range nums[1:] {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		fmt.Println("sum:", sum, "max:", max)
		max = myMax(max, sum)
	}
	return max
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
