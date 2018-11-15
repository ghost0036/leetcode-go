package leetCode_1

func twoSum(nums []int, target int) []int {
	for index, v := range nums {
		for index1, v1 := range nums[index+1:] {
			if v+v1 == target {
				return []int{index, index1 + index + 1}
			}
		}
	}
	return []int{}
}
