package leetcode

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) (result string) {
	m := make(map[string]bool)
	for i := range words {
		if _, ok := m[words[i]]; !ok {
			m[words[i]] = true
		}
	}

	for _, w := range words {
		if len(w) >= len(result) {
			isVaild := true

			if len(w) == len(result) {
				for i := range w {
					if w[i] > result[i] {
						isVaild = false
						break
					} else if w[i] < result[i] {
						break
					}
				}
			}

			if isVaild {
				for i := 0; i < len(w); i++ {
					if i == len(w)-1 {
						result = w
						break
					}
					if _, ok := m[string(w[:i+1])]; !ok {
						break
					}
				}
			}
		}
	}

	return
}

// @lc code=end
