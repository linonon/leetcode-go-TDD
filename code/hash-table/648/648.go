package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dictionary []string, sentence string) string {
	bronkenString := strings.Fields(sentence)
	var result []string
	var dicVal string

	for _, val := range bronkenString {
		var tempRes string
		for _, dicVal = range dictionary {
			if strings.HasPrefix(val, dicVal) {
				if len(tempRes) == 0 {
					tempRes = dicVal
				} else if len(dicVal) < len(tempRes) {
					tempRes = dicVal
				}
			}
		}

		if len(tempRes) > 0 {
			result = append(result, tempRes)
		} else {
			result = append(result, val)
		}
	}

	return strings.Join(result, " ")
}

// @lc code=end
