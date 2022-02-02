package leetcode

/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] 地图中的最高点
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	type pair struct{ x, y int }
	// 多元方向
	var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	m, n := len(isWater), len(isWater[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		for j := range result[i] {
			// 初始化未探索的点 -1
			result[i][j] = -1
		}
	}

	// 水源入队
	q := []pair{}
	for i, row := range isWater {
		for j, water := range row {
			if water == 1 {
				result[i][j] = 0
				q = append(q, pair{i, j})
			}
		}
	}

	for len(q) > 0 {
		// 探索顺序： 水源 -> 未探索的入队点
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			x := p.x + d.x
			y := p.y + d.y

			if 0 <= x && x < m && 0 <= y && y < n && result[x][y] == -1 {
				result[x][y] = result[p.x][p.y] + 1
				// 添加需要探索的点
				q = append(q, pair{x, y})
			}
		}
	}

	return result
}

// @lc code=end
