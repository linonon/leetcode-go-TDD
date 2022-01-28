# 比赛中的配对次数

## Version 1
按著邏輯走
```go
func numberOfMatches(n int) int {
	result := 0

	for n != 1 {
		if n%2 == 1 {
			n--
			n = n / 2
			result += n
			n++
		} else {
			n = n / 2
			result += n
		}
	}

	return result
}
```

## Version 2 (Math)
哈哈，因為就是要從 n 個人選一個王者出來， 所要要進行 n-1場比賽。
```go
func numberOfMatches(n int) int {
    return n - 1
}
```