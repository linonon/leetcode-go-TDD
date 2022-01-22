# 加一

## Version1

```go
func plusOne(digits []int) []int {
	sum := 0
	for i := len(digits) - 1; i >= 0; i-- {
		timer := 1
		for j := i - 1; j >= 0; j-- {
			timer *= 10
		}

		sum += timer * digits[len(digits)-1-i]
		fmt.Println("sum + 1 =", sum)
	}

	sum += 1
	stringSum := strconv.Itoa(sum)
	result := make([]int, len(stringSum))
	for i := 0; i < len(stringSum); i++ {
		timer := 1
		for j := len(stringSum) - 1 - i; j > 0; j-- {
			timer *= 10
		}
		if i == 0 {
			result[i] = sum / timer
		} else if i > 0 && i < len(stringSum)-1 {
			result[i] = (sum / timer) % 10
		} else {
			result[i] = sum % 10
		}
	}

	return result
}
```

這個方法不行：當 len([]int) >> int64 時，就會出錯。

## Version2
用Recrusion的方法做，完美過關
```go
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
```

## Version "Most Vote"

```go
func plusOne(digits []int) []int {
    var n int = len(digits)
    for i:= n-1; i>=0; i--{
        if digits[i] < 9 {
            digits[i]+=1
            return digits
        } else {
            digits[i] = 0
        }
    }
    var a = make([]int,n+1)
    a[0] = 1
    return a

}
```