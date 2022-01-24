# 實現 strStr()

## Version 1
暴力法
```go
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	length := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}

	return -1
}
```


## KMP
- pi(i)[0 <= i < len(s)]: `s[0 : i]` 內，最長的重複字串長度
- 兩個重要的性質
    1. pi(i) <= pi(i-1) + 1
    2. 如果 `s[0:pi(i)-1] = s[i-(pi(i)-1):i]`, 則 `pi(i) = p(i-1)+1`

```go
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && needle[i] != needle[j] {
			j = pi[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = pi[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}
```