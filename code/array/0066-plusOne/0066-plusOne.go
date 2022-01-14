package leetcode

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	tPlusOne(digits)

	if digits[0] == 10 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}

func tPlusOne(digits []int) {
	digits[len(digits)-1] += 1

	if digits[len(digits)-1] == 10 && len(digits) > 1 {
		digits[len(digits)-1] = 0
		tPlusOne(digits[:len(digits)-1])
	}
}

// @lc code=end
