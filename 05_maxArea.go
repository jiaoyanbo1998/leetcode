package leetcode

/*
	给定一个长度为 n 的整数数组 height 。
	有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])

	输入：[1,8,6,2,5,4,8,3,7]
	输出：49

	双指针：left == 0 right == len(height) - 1
		   左指针向右移动
		   右指针向左移动
	底 == right - left
	高 == min(height[left], height[right])
	面积 == max(面积, 底 * 高)
	如果 左边高度 < 右边高度
			左指针向右移动
	否则 右指针向左移动

	本质：双指针，左比右小，左指针右移，否则右指针左移，直到相遇，返回最大面积
*/
func maxArea(height []int) (area int) {
	left := 0                // 左指针
	right := len(height) - 1 // 右指针
	for right > left {
		// 底
		base := right - left
		// 高
		high := min(height[left], height[right])
		// 面积
		area = max(area, base*high)
		// 左指针右移
		if height[left] < height[right] {
			left++
		} else { // 右指针左移
			right--
		}
	}
	return area
}
