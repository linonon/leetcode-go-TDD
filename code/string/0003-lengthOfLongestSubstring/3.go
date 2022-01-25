package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 判定最基礎的情況
	if len(s) <= 1 {
		return len(s)
	}

	var result int
	// flag 用來調整 j，區間 [j：i] 是滑動窗口
	flag := 0
	tmp := make([]byte, 1)
	for i := 1; i < len(s); i++ {
		tmp[0] = s[0]
		for j := flag; j < i; j++ {
			if s[i] != s[j] {
				tmp = append(tmp, s[i])
			} else {
				// 如果發現相等，調整滑動窗口
				flag = j + 1
				j = i
			}

			// 判斷 result 是否是 max
			if result < len(tmp) {
				result = len(tmp)
			}
		}
		// 類似初始化操作，Cap信息仍保存
		tmp = tmp[0:1]
	}

	return result
}

// @lc code=end
