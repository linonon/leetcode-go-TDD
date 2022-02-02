package leetcode

/*
 * @lc app=leetcode.cn id=2000 lang=golang
 *
 * [2000] 反转单词前缀
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	var result string
	for i := range word {
		if word[i] == ch {
			for j := i; j >= 0; j-- {
				result += string(word[j])
			}
			return result + word[i+1:]
		}
	}
	return word
}

// @lc code=end
