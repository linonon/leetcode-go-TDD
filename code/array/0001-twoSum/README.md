# 兩數之和

這題沒什麼好說的吧
```go 
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		j, ok := m[-v]
		m[v-target] = i
		if ok {
			return []int{j, i}
		}
	}

	return []int{}
}
```