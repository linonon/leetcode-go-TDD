package leetcode

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var reversedNum int
	for x > reversedNum {
		reversedNum = reversedNum*10 + x%10
		x = x / 10
	}

	return x == reversedNum || x == reversedNum/10
}

// @lc code=end
