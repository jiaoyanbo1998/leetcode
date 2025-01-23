package leetcode

/*
	给定一个整数数组nums和一个整数目标值target，
	请你在该数组中找出和为目标值target的两个整数，返回它们的数组下标。

	输入：nums = [2,7,11,15], target = 9
	输出：[0,1]

	创建一个map，key为数据，value为下标
	遍历nums
	判断target-value是否存在于map中
	不在，将key == value，value == index存入map中
	存在，返回[]int{map[target-value], index}

	本质就是，将数据存在map中，然后target挨个减去nums的元素，看是否在map中
*/
// twoSum 两数之和
func twoSum(nums []int, target int) []int {
	// 创建一个map，key == 值，value == 下标
	m := make(map[int]int, len(nums))
	for index, value := range nums {
		// 判断[target - value]是否在map中
		res, ok := m[target-value]
		// 在map中
		if ok {
			return []int{res, index}
		}
		// 不在map中
		m[value] = index
	}
	return nil
}
