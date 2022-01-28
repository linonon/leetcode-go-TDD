package leetcode

/*
 * @lc app=leetcode.cn id=1688 lang=golang
 *
 * [1688] 比赛中的配对次数
 */

// @lc code=start
func numberOfMatches(n int) int {
	result := 0

	for n != 1 {
		if n%2 == 1 {
			n--
			n = n / 2
			result += n
			n++
		} else {
			n = n / 2
			result += n
		}
	}

	return result
}

// @lc code=end
