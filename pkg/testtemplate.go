package pkgu

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T, flag interface{}, fn func(t *testing.T)) {
	var str string
	if flag == nil {
		str = "Round"
	}

	num, ok := flag.(int)
	if ok {
		str = fmt.Sprintf("Round-%v", num)
	}

	text, ok := flag.(string)
	if ok {
		str = fmt.Sprintf("%v\n", text)
	}
	t.Run(str, fn)
}
