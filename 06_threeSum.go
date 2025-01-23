package leetcode

import "sort"

/*
	给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
	满足 i != j、i != k 且 j != k
	满足 nums[i] + nums[j] + nums[k] == 0
	请你返回所有和为 0 且不重复的三元组

	注意：答案中不可以包含重复的三元组

	输入：nums = [-1,0,1,2,-1,-4]
	输出：[[-1,-1,2],[-1,0,1]]

	本质：外层循环+双指针
*/

// threeSum 三数之和
func threeSum(nums []int) (res [][]int) {
	// 排序
	sort.Ints(nums)
	// 外层循环
	for i := 0; i < len(nums); i++ {
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针
		left, right := i+1, len(nums)-1
		for left < right {
			// 去重
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}
			// 计算三数之和
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}
