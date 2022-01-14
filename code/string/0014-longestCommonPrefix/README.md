# 最長公共前綴

## TestData
```go
{
    TestData: []string{"flower", "flow", "flight"},
    Want:     "fl",
},
{
    TestData: []string{"dog", "racecar", "car"},
    Want:     "",
},
{
    TestData: []string{"dogggg", "dog", "do"},
    Want:     "do",
},
```

## Version 1
```go
func longestCommonPrefix(strs []string) string {
	var result = ""
	for i := 0; i < len(strs[0]); i++ {
		checker := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return result
			}
			if checker != string(strs[j][i]) {
				return result
			}
		}
		result += checker
	}

	return result
}
```