package leetcode

/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 */

// @lc code=start
func hasAlternatingBits(n int) (has bool) {
	if n <= 2 {
		return true
	}

	var mod int
	prevmod := -1
	for n != 0 {
		mod = n % 2
		if prevmod == mod {
			return false
		}
		n = n / 2
		prevmod = mod
		if n == 1 && n == prevmod {
			return false
		}
	}

	return true
}

// @lc code=end
