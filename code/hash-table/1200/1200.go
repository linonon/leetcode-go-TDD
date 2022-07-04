package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1200 lang=golang
 *
 * [1200] 最小绝对差
 */

// @lc code=start
func minimumAbsDifference(arr []int) (result [][]int) {
	sort.Ints(arr)
	minDiff := arr[1] - arr[0]

	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minDiff {
			minDiff = arr[i+1] - arr[i]
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}
	return
}

// @lc code=end
