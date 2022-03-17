package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start
func longestWord(words []string) (result string) {
	m := make(map[int][]string)
	max := 0
	for i := range words {
		m[len(words[i])] = append(m[len(words[i])], string(words[i]))
		if len(words[i]) > max {
			max = len(words[i])
		}
	}
	fmt.Println(m)

	for i := max; i > 0; i-- {
		for _, world := range m[i] {
			if _, ok := m[i]; !ok {
				continue
			}
		}
	}
	return
}

// @lc code=end
