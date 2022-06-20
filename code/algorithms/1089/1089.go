package leetcode

/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start
func duplicateZeros(arr []int) {
	zero := 0
	for _, i := range arr {
		if i == 0 {
			zero++
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if zero+i < len(arr) {
				arr[zero+i] = 0
			}
			if zero-1+i < len(arr) {
				arr[zero-1+i] = 0
			}
			zero--
		} else {
			if zero+i < len(arr) {
				arr[zero+i] = arr[i]
			}
		}
	}
}

// @lc code=end
