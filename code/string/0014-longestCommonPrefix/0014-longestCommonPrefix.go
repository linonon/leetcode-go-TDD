package leetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	var result = ""
	for i := 0; i < len(strs[0]); i++ {
		checker := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) {
				return result
			}
			if checker != string(strs[j][i]) {
				return result
			}
		}
		result += checker
	}

	return result
}

// @lc code=end
