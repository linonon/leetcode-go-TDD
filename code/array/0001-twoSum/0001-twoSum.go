package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		j, ok := m[-v]
		m[v-target] = i
		if ok {
			return []int{j, i}
		}
	}

	return []int{}
}

// @lc code=end
