package leetcode

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	romanMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := romanMap[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[string(s[i])] >= romanMap[string(s[i+1])] {
			result += romanMap[string(s[i])]
		} else {
			result -= romanMap[string(s[i])]
		}
	}

	return result
}

// @lc code=end
