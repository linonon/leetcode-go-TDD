package leetcode

/*
 * @lc app=leetcode.cn id=2049 lang=golang
 *
 * [2049] 统计最高分的节点数目
 */

// @lc code=start
func countHighestScoreNodes(parents []int) (ans int) {
	n := len(parents)
	// 每個點的子節點
	children := make([][]int, n)
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}

	maxScore := 0
	var dfs func(int) int
	dfs = func(node int) int {
		score, size := 1, n-1
		for _, ch := range children[node] {
			sz := dfs(ch)
			score *= sz
			size -= sz
		}
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			maxScore = score
			ans = 1
		}
		return n - size
	}

	dfs(0)
	return
}

// @lc code=end
