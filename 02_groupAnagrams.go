package leetcode

import "slices"

/*
	给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表

	输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
	输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

	创建一个map，key为string，value为[]string
	遍历strs
	比如：将"eat"转为['e', 'a', 't']
         ['e', 'a', 't']升序排序为['a', 'e', 't']
		 将['a', 'e', 't']转为"aet"
		 将"eat"插入map，key == "aet"，value == "eat"
	     -------------------------------------------
		 将"ate"转为['a', 't', 'e]
		 ['a', 't', 'e']升序排序为['a', 'e', 't']
         将['a', 'e', 't']转为"aet"
		 将"ate"插入map，key == "aet"，value == "eat", "ate"
	     --------------------------------------------------
					......
	将map的值转为二维数组

	本质：将乱序的字符串，存储到同一个key(排好序的key)下
*/

// groupAnagrams 字母异位词分组
func groupAnagrams(strs []string) [][]string {
	// 创建map
	m := make(map[string][]string, len(strs))
	// 遍历strs
	for _, str := range strs {
		// 将str转为字符切片
		bytes := []byte(str)
		// 字符切片从小到大排序
		slices.Sort(bytes)
		// 将字符数组转为字符串
		s := string(bytes)
		// 将str插入map，key为排好序的字符串，所有的乱序都存储在同一个key下
		m[s] = append(m[s], str)
	}
	// 将map转为二维切片
	ss := make([][]string, 0, len(m))
	// value为切片
	for _, value := range m {
		ss = append(ss, value)
	}
	return ss
}
