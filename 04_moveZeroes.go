package leetcode

/*
	给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
		输入: nums = [0,1,0,3,12]
		输出: [1,3,12,0,0]

	双指针：right = -1， left = 0
	right向又移动，直到len(nums) - 1
	如果nums[right] == 0，跳过，继续下次循环
	如果nums[right] != 0，交换nums[right]和nums[left]，left++

	本质：双指针，遇见非0元素，交换左右指针的值
*/
func moveZeroes(nums []int) {
	left := 0                                    // 左指针
	for right := 0; right < len(nums); right++ { // right 右指针
		// 元素为0，跳过循环
		if nums[right] == 0 {
			continue
		}
		// 左右指针不指向同一个位置，交换左右指针的值
		if left != right {
			nums[left], nums[right] = nums[right], nums[left]
		}
		// 左指针右移
		left++
	}
}
