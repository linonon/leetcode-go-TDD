package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=2047 lang=golang
 *
 * [2047] 句子中的有效单词数
 */

// @lc code=start
func countValidWords(s string) int {
	m := map[rune]bool{
		'0': true, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true,
		'!': true, '.': true, ',': true,
	}
	result := 0
	for _, tmp := range strings.Fields(s) {

		if len(tmp) > 1 && (tmp[len(tmp)-1] == '.' || tmp[len(tmp)-1] == '!' || tmp[len(tmp)-1] == ',') && !m[rune(tmp[len(tmp)-2])] {
			tmp = tmp[:len(tmp)-1]
		}

		isGoodWord := true
		count := 0
		for i := 0; i < len(tmp); i++ {
			if m[rune(tmp[i])] || tmp[0] == '-' || tmp[len(tmp)-1] == '-' {
				isGoodWord = false
				break
			}
			if tmp[i] == '-' {
				count++
				if count >= 2 {
					isGoodWord = false
				}
			}
		}

		if len(tmp) == 1 && (tmp[0] == '!' || tmp[0] == ',' || tmp[0] == '.') {
			isGoodWord = true
		}

		if isGoodWord && len(tmp) > 0 {
			result++
		}
	}

	return result
}

// @lc code=end
