package leetcode

/*
	给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
		输入：nums = [100,4,200,1,3,2]
		输出：4
		解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

	创建一个map，key为int，value为bool
	将数组转成hash
		map[int]bool = {
			100:true,
			4:true,
			200:true,
			1:true,
			3:true,
			2:true,
		}
	遍历hash
	比如：for v := range m {
			v == 100，m[100-1] == m[99] == false，100是起始数字
         	start == 100
		 	for map[start] { start存在，start++
		 		start++
		 	}
		 	res = max(res, start - v)
		}

	本质：是一个滑动窗口，不停的扩大窗口，直到窗口内的数字不连续为止
*/
func longestConsecutive(nums []int) (res int) {
	// 创建map
	m := make(map[int]bool, len(nums))
	// 将切片转为map
	for _, val := range nums {
		m[val] = true // true，在窗口，false，不在窗口
	}
	// 遍历map
	for key := range m {
		// 判断当前数字是否是，起始数字
		if m[key-1] == true {
			continue // 不是起始数字，直接跳过本次循环
		}
		// 是起始数字，扩大窗口
		start := key
		for m[start] == true {
			start++ // map中，有的key都为true，中断了会变为false，结束for循环
		}
		// 获取最大值
		res = max(res, start-key)
	}
	return res
}
