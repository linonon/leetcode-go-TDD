package leetcode

/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] 统计最高分的节点数目
 */

// @lc code=start
func countHighestScoreNodes(parents []int) int {
	var m = make(map[int][]int)

	for i := range parents {
		for j := range parents {
			if parents[j] == i {
				m[j] = append(m[j], j)
			}
		}
	}

	return 0
}

// @lc code=end
