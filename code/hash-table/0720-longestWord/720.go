package leetcode

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) (result string) {
	m := make(map[string]int)
	for i := range words {
		if _, ok := m[words[i]]; !ok {
			m[words[i]] = 1
		}
	}

	result = ""
	for _, w := range words {
		isVaild := true

		if len(w) >= len(result) {
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
				for i := 0; i < len(w)-1; i++ {
					s := string(w[:i+1])
					if _, ok := m[s]; !ok {
						isVaild = false
						break
					}
				}
			}
		} else {
			isVaild = false
		}

		if isVaild {
			result = w
		}
	}

	return
}

// @lc code=end
