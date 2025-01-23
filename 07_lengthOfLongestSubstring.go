package leetcode

/*
	给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度

	输入: s = "abcabcbb"
	输出: 3
	解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

	输入: s = "bbbbb"
	输出: 1
	解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

	本质：滑动窗口，true在窗口内，false不在窗口内
*/
func lengthOfLongestSubstring(s string) (res int) {
	// 创建窗口，长度为128的bool数组
	window := [128]bool{}
	left := 0 // 左指针
	// 遍历字符串
	for right, val := range s {
		// 当前字符在窗口内
		for window[val] == true {
			// 将值设置为false
			window[s[left]] = false
			// 窗口右移
			left++
		}
		// 将当前字符插入窗口
		window[val] = true
		// 获取最大窗口长度
		res = max(res, right-left+1) // +1是因为索引从0开始
	}
	return res
}
