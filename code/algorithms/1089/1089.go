package leetcode

/*
 * @lc app=leetcode.cn id=1089 lang=golang
 *
 * [1089] 复写零
 */

// @lc code=start
func duplicateZeros(arr []int) {
	n := len(arr)
	var tail []int
	for i := 0; i < n; i++ {
		if len(tail)-2 >= 0 {
			for j := 0; j < len(tail)-1; j++ {
				if j+i == n {
					break
				}
				arr[j+i] = tail[j]
			}
			tail = tail[:0]
		}
		if arr[i] == 0 {
			if i+1 == n {
				break
			}
			tail = append([]int{}, arr[i+1:]...)
			arr[i+1] = 0
			i++
		}
	}
}

// @lc code=end
