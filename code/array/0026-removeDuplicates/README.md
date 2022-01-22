# 删除有序数组中的重复项

## Version 1
Easy, But Slowly

原因是因為沒讀懂題目！

LeetCode 是根據我們結果 int 去用 循環前 i 個Slice。所以我沒必要對輸入的 Slice 進行改動

```go
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 0; i < len(nums); {
		if i < len(nums)-1 && nums[i] == nums[i+1] {
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}

		i++
	}

	return len(nums)
}
```

## Version Most Votes
```go
func removeDuplicates(nums []int) int {
    ln := len(nums)
    if ln <= 1 {
        return ln
    }
    
    j := 0 // points to  the index of last filled position
    for i := 1; i < ln; i++ {
        if nums[j] != nums[i] {
            j++
            nums[j] = nums[i]
        }
    }
    
    return j + 1
}
```