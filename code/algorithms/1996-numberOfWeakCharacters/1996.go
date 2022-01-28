package leetcode

import "sort"

// import "sync"

/*
 * @lc app=leetcode.cn id=1996 lang=golang
 *
 * [1996] 游戏中弱角色的数量
 */

// @lc code=start
func numberOfWeakCharacters(properties [][]int) int {
	result := 0
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] > q[0] || (p[0] == q[0] && p[1] < q[1])
	})
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			result++
		} else {
			maxDef = p[1]
		}
	}
	return result
}

// @lc code=end
