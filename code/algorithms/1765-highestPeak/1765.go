package leetcode

import "reflect"

/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}

	return result
}

// @lc code=end
