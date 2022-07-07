package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dictionary []string, sentence string) string {
	ss := strings.Split(sentence, " ")
	sort.Slice(dictionary, func(i, j int) bool {
		if len(dictionary[i]) < len(dictionary[j]) {
			return true
		}
		return false
	})

	fmt.Println(dictionary)
	for i, s := range ss {
		for _, keyword := range dictionary {
			if strings.Contains(s, keyword) {
				ss[i] = keyword
				break
			}
		}
	}

	return strings.Join(ss, " ")
}

// @lc code=end
